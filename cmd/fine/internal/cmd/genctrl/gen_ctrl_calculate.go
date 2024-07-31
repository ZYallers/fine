package genctrl

import (
	"github.com/ZYallers/fine/cmd/fine/internal/consts"
	"github.com/ZYallers/fine/cmd/fine/internal/util/utils"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/text/fregex"
	"github.com/ZYallers/fine/text/fstr"
)

func (g *genCtrl) getApiItemsInSrc(apiModuleFolderPath string) (items []apiItem, err error) {
	var importPath string
	// The second level folders: versions.
	apiVersionFolderPaths, err := ffile.ScanDir(apiModuleFolderPath, "*", false)
	if err != nil {
		return nil, err
	}
	for _, apiVersionFolderPath := range apiVersionFolderPaths {
		if !ffile.IsDir(apiVersionFolderPath) {
			continue
		}
		// The second level folders: versions.
		apiFileFolderPaths, err := ffile.ScanDir(apiVersionFolderPath, "*.go", false)
		if err != nil {
			return nil, err
		}
		importPath = utils.GetImportPath(apiVersionFolderPath)
		for _, apiFileFolderPath := range apiFileFolderPaths {
			if ffile.IsDir(apiFileFolderPath) {
				continue
			}
			structsInfo, err := g.getStructsNameInSrc(apiFileFolderPath)
			if err != nil {
				return nil, err
			}
			for _, methodName := range structsInfo {
				// remove end "Req"
				methodName = fstr.TrimRightStr(methodName, "Req", 1)
				item := apiItem{
					Import:     fstr.Trim(importPath, `"`),
					FileName:   ffile.Name(apiFileFolderPath),
					Module:     ffile.Basename(apiModuleFolderPath),
					Version:    ffile.Basename(apiVersionFolderPath),
					MethodName: methodName,
				}
				items = append(items, item)
			}
		}
	}
	return
}

func (g *genCtrl) getApiItemsInDst(dstFolder string) (items []apiItem, err error) {
	if !ffile.Exists(dstFolder) {
		return nil, nil
	}
	type importItem struct {
		Path  string
		Alias string
	}
	filePaths, err := ffile.ScanDir(dstFolder, "*.go", true)
	if err != nil {
		return nil, err
	}
	for _, filePath := range filePaths {
		var (
			array       []string
			importItems []importItem
			importLines []string
			module      = ffile.Basename(ffile.Dir(filePath))
		)
		importLines, err = g.getImportsInDst(filePath)
		if err != nil {
			return nil, err
		}

		// retrieve all imports.
		for _, importLine := range importLines {
			array = fstr.SplitAndTrim(importLine, " ")
			if len(array) == 2 {
				importItems = append(importItems, importItem{
					Path:  fstr.Trim(array[1], `"`),
					Alias: array[0],
				})
			} else {
				importItems = append(importItems, importItem{
					Path: fstr.Trim(array[0], `"`),
				})
			}
		}
		// retrieve all api usages.
		// retrieve it without using AST, but use regular expressions to retrieve.
		// It's because the api definition is simple and regular.
		// Use regular expressions to get better performance.
		fileContent := ffile.GetContents(filePath)
		matches, err := fregex.MatchAllString(consts.PatternCtrlDefinition, fileContent)
		if err != nil {
			return nil, err
		}
		for _, match := range matches {
			// try to find the import path of the api.
			var (
				importPath string
				version    = match[1]
				methodName = match[2] // not the function name, but the method name in api definition.
			)
			for _, item := range importItems {
				if item.Alias != "" {
					if item.Alias == version {
						importPath = item.Path
						break
					}
					continue
				}
				if ffile.Basename(item.Path) == version {
					importPath = item.Path
					break
				}
			}
			item := apiItem{
				Import:     fstr.Trim(importPath, `"`),
				Module:     module,
				Version:    ffile.Basename(importPath),
				MethodName: methodName,
			}
			items = append(items, item)
		}
	}
	return
}
