package genservice

import (
	"fmt"
	"log"
	"strings"

	"github.com/ZYallers/fine/cmd/fine/internal/util/utils"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/text/fregex"
	"github.com/ZYallers/fine/text/fstr"
)

type packageItem struct {
	Alias     string
	Path      string
	RawImport string
}

type genService struct {
	srcFolder string
	dstFolder string
}

func (g *genService) Run() error {
	if !ffile.Exists(g.srcFolder) {
		return fmt.Errorf("source folder path \"%s\" not exist", g.srcFolder)
	}
	if !ffile.Exists(g.dstFolder) {
		return fmt.Errorf("destination folder \"%s\" not exist", g.dstFolder)
	}

	importPrefix := utils.GetImportPath(g.srcFolder)

	// The first level folders.
	srcFolderPaths, err := ffile.ScanDir(g.srcFolder, "*", false)
	if err != nil {
		return err
	}

	var (
		isDirty               bool     // Temp boolean.
		initImportSrcPackages []string // Used for generating logic.go.
	)

	for _, srcFolderPath := range srcFolderPaths {
		if !ffile.IsDir(srcFolderPath) {
			continue
		}
		// Only retrieve sub files, no recursively.
		files, err := ffile.ScanDir(srcFolderPath, "*.go", false)
		if err != nil {
			return err
		}
		if len(files) == 0 {
			continue
		}

		// Parse single logic package folder.
		var (
			srcPackageName       = ffile.Basename(srcFolderPath)
			dstFilePath          = ffile.Join(g.dstFolder, g.getDstFileNameCase(srcPackageName, "Snake")+".go")
			srcCodeCommentedMap  = make(map[string]string)
			importAliasToPathMap = make(map[string]string) // for conflict imports check. alias => import path(with `"`)
			importPathToAliasMap = make(map[string]string) // for conflict imports check. import path(with `"`) => alias
			srcImportedPackages  = make([]string, 0)
			srcPkgInterfaceMap   = make(map[string][]string)
		)

		for _, file := range files {
			fileContent := ffile.GetContents(file)
			// Calculate code comments in source Go files.
			err := g.calculateCodeCommented(fileContent, srcCodeCommentedMap)
			if err != nil {
				return err
			}

			// remove all comments.
			fileContent, err = fregex.ReplaceString(`(//.*)|((?s)/\*.*?\*/)`, "", fileContent)
			if err != nil {
				return err
			}

			// Calculate imported packages of source go files.
			packageItems, err := g.calculateImportedPackages(fileContent)
			if err != nil {
				return err
			}

			// try finding the conflicts imports between files.
			for _, item := range packageItems {
				var alias = item.Alias
				if alias == "" {
					alias = ffile.Basename(fstr.Trim(item.Path, `"`))
				}

				// ignore unused import paths, which do not exist in function definitions.
				if !fregex.IsMatchString(fmt.Sprintf(`func .+?([^\w])%s(\.\w+).+?{`, alias), fileContent) {
					continue
				}

				// find the exist alias with the same import path.
				existAlias, ok := importPathToAliasMap[item.Path]
				if ok && existAlias != "" {
					fileContent, err = fregex.ReplaceStringFuncMatch(
						fmt.Sprintf(`([^\w])%s(\.\w+)`, alias),
						fileContent,
						func(match []string) string {
							return match[1] + existAlias + match[2]
						},
					)
					if err != nil {
						return err
					}
					continue
				}

				// resolve alias conflicts.
				importPath, _ := importAliasToPathMap[alias]
				if importPath == "" {
					importAliasToPathMap[alias] = item.Path
					importPathToAliasMap[item.Path] = alias
					srcImportedPackages = append(srcImportedPackages, item.RawImport)
					continue
				}
				if importPath != item.Path {
					// update the conflicted alias for import path with suffix.
					// eg:
					// v1  -> v10
					// v11 -> v110
					for aliasIndex := 0; ; aliasIndex++ {
						item.Alias = fmt.Sprintf(`%s%d`, alias, aliasIndex)
						existPathForAlias, _ := importAliasToPathMap[item.Alias]
						if existPathForAlias != "" {
							if existPathForAlias == item.Path {
								break
							}
							continue
						}
						break
					}
					importPathToAliasMap[item.Path] = item.Alias
					importAliasToPathMap[item.Alias] = item.Path
					// reformat the import path with alias.
					item.RawImport = fmt.Sprintf(`%s %s`, item.Alias, item.Path)
					// update the file content with new alias import.
					fileContent, err = fregex.ReplaceStringFuncMatch(
						fmt.Sprintf(`([^\w])%s(\.\w+)`, alias),
						fileContent,
						func(match []string) string {
							return match[1] + item.Alias + match[2]
						},
					)
					if err != nil {
						return err
					}
					srcImportedPackages = append(srcImportedPackages, item.RawImport)
				}
			}

			srcImportedPackages = utils.RemoveDuplicateWithString(srcImportedPackages)

			// Calculate functions and interfaces for service generating.
			err = g.calculateInterfaceFunctions(fileContent, srcPkgInterfaceMap)
			if err != nil {
				return err
			}
		}

		initImportSrcPackages = append(initImportSrcPackages, fmt.Sprintf(`%s/%s`, importPrefix, srcPackageName))

		// Generating service go file for single logic package.
		ok, err := g.generateServiceFile(srcImportedPackages, srcPkgInterfaceMap, srcCodeCommentedMap, dstFilePath)
		if err != nil {
			return err
		}
		if ok {
			isDirty = true
		}
	}

	if isDirty {
		// Generate initialization go file.
		if len(initImportSrcPackages) > 0 {
			if err = g.generateInitializationFile(initImportSrcPackages); err != nil {
				return err
			}
		}
		utils.GoFmt(g.dstFolder)
	}

	// auto update main.go.
	if err := g.checkAndUpdateMain(); err != nil {
		return err
	}

	return nil
}

// getDstFileNameCase call str.Case* function to convert the str to specified case.
func (g *genService) getDstFileNameCase(str, caseStr string) (newStr string) {
	if newStr := fstr.CaseConvert(str, fstr.CaseTypeMatch(caseStr)); newStr != str {
		return newStr
	}
	return fstr.CaseSnake(str)
}

// insertBefore inserts the `values` to the front of `index`.
func (g *genService) insertBefore(arr []string, index int, values ...string) error {
	if index < 0 || index >= len(arr) {
		return fmt.Errorf("index %d out of array range %d", index, len(arr))
	}
	rear := append([]string{}, arr[index:]...)
	arr = append(arr[0:index], values...)
	arr = append(arr, rear...)
	return nil
}

func (g *genService) checkAndUpdateMain() (err error) {
	var (
		logicPackageName = fstr.ToLower(ffile.Basename(g.srcFolder))
		logicFilePath    = ffile.Join(g.srcFolder, logicPackageName+".go")
		importPath       = utils.GetImportPath(logicFilePath)
		importStr        = fmt.Sprintf(`_ "%s"`, importPath)
		mainFilePath     = ffile.Join(ffile.Dir(ffile.Dir(ffile.Dir(logicFilePath))), "main.go")
		mainFileContent  = ffile.GetContents(mainFilePath)
	)
	// No main content found.
	if mainFileContent == "" {
		return nil
	}
	if fstr.Contains(mainFileContent, importStr) {
		return nil
	}
	match, err := fregex.MatchString(`import \(([\s\S]+?)\)`, mainFileContent)
	if err != nil {
		return err
	}
	// No match.
	if len(match) < 2 {
		return nil
	}
	lines := fstr.Split(match[1], "\n")
	for i, line := range lines {
		line = fstr.Trim(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == '_' {
			continue
		}
		// Insert the logic import into imports.
		if err = g.insertBefore(lines, i, fmt.Sprintf("\t%s\n\n", importStr)); err != nil {
			return err
		}
		break
	}
	mainFileContent, err = fregex.ReplaceString(
		`import \(([\s\S]+?)\)`,
		fmt.Sprintf(`import (%s)`, strings.Join(lines, "\n")),
		mainFileContent,
	)
	if err != nil {
		return err
	}
	log.Printf("update main.go file: %s\n", mainFilePath)
	err = ffile.PutContents(mainFilePath, mainFileContent)
	if err != nil {
		return fmt.Errorf("update main.go file failed: %s", err)
	}
	utils.GoFmt(mainFilePath)
	return
}
