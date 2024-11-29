package genservice

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/ZYallers/fine/cmd/fine/internal/consts"
	"github.com/ZYallers/fine/cmd/fine/internal/util/utils"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/text/fregex"
	"github.com/ZYallers/fine/text/fstr"
)

func (g *genService) generateServiceFile(srcImportedPackages []string, srcStructFunctions map[string][]string,
	srcCodeCommentedMap map[string]string, dstFilePath string) (ok bool, err error) {
	var (
		generatedContent        string
		allFuncArray            = make([]string, 0) // Used for check whether interface dirty, going to change file content.
		importedPackagesContent = fmt.Sprintf("import (\n%s\n)", fstr.Join(srcImportedPackages, "\n"))
		dstFileName             = fstr.CaseCamel(ffile.Name(dstFilePath))
	)

	dstPackageName := fstr.ToLower(ffile.Basename(g.dstFolder))
	generatedContent += fstr.ReplaceByMap(consts.TemplateGenServiceContentHead, map[string]string{
		"{Imports}":     importedPackagesContent,
		"{PackageName}": dstPackageName,
		"{CreatedAt}":   utils.CreatedAt(),
	})

	// Type definitions.
	generatedContent += "type("
	generatedContent += "\n"
	for key, value := range srcStructFunctions {
		structName, funcArray := key, value
		allFuncArray = append(allFuncArray, funcArray...)
		// Add comments to a method.
		for index, funcName := range funcArray {
			sfn := fmt.Sprintf("%s-%s", structName, funcName)
			if commentedInfo, exist := srcCodeCommentedMap[sfn]; exist {
				if str := fstr.Trim(commentedInfo); str != "" {
					funcName = str + "\n" + funcName
					// if commentedArray := fstr.Split(str, "\n"); len(commentedArray) > 0 {
					// 	funcName += " " + fstr.Trim(commentedArray[0])
					// }
				}
				funcArray[index] = funcName
			}
		}
		generatedContent += fstr.Trim(fstr.ReplaceByMap(consts.TemplateGenServiceContentInterface, map[string]string{
			"{InterfaceName}":  "I" + dstFileName + structName,
			"{FuncDefinition}": strings.Join(funcArray, "\n\t"),
		}))
		generatedContent += "\n"
	}
	generatedContent += ")"
	generatedContent += "\n"

	// Generating variable and register definitions.
	var (
		variableContent          string
		generatingInterfaceCheck string
	)
	// Variable definitions.
	for key, _ := range srcStructFunctions {
		structName := key
		generatingInterfaceCheck = fmt.Sprintf(`[^\w\d]+%s.I%s[^\w\d]`, dstPackageName, structName)
		if fregex.IsMatchString(generatingInterfaceCheck, generatedContent) {
			continue
		}
		variableContent += fstr.Trim(fstr.ReplaceByMap(consts.TemplateGenServiceContentVariable, map[string]string{
			"{StructName}":    dstFileName + structName,
			"{InterfaceName}": "I" + dstFileName + structName,
		}))
		variableContent += "\n"
	}
	if variableContent != "" {
		generatedContent += "var("
		generatedContent += "\n"
		generatedContent += variableContent
		generatedContent += ")"
		generatedContent += "\n"
	}
	// Variable register function definitions.
	for key, _ := range srcStructFunctions {
		structName := key
		generatingInterfaceCheck = fmt.Sprintf(`[^\w\d]+%s.I%s[^\w\d]`, dstPackageName, structName)
		if fregex.IsMatchString(generatingInterfaceCheck, generatedContent) {
			continue
		}
		generatedContent += fstr.Trim(fstr.ReplaceByMap(consts.TemplateGenServiceContentRegister, map[string]string{
			"{StructName}":    dstFileName + structName,
			"{InterfaceName}": "I" + dstFileName + structName,
		}))
		generatedContent += "\n\n"
	}

	// Replace empty braces that have new line.
	generatedContent, _ = fregex.ReplaceString(`{[\s\t]+}`, `{}`, generatedContent)

	// Remove package name calls of `dstPackageName` in produced codes.
	generatedContent, _ = fregex.ReplaceString(fmt.Sprintf(`\*{0,1}%s\.`, dstPackageName), ``, generatedContent)

	// Write file content to disk.
	if ffile.Exists(dstFilePath) {
		if !utils.IsFileDoNotEdit(dstFilePath) {
			log.Printf("ignore file as it is manually maintained: %s\n", dstFilePath)
			return false, nil
		}
		if !g.isToGenerateServiceGoFile(dstPackageName, dstFilePath, allFuncArray) {
			log.Printf("not dirty, ignore generate: %s\n", ffile.RealPath(dstFilePath))
			return false, nil
		}
	}
	if err = ffile.PutContents(dstFilePath, generatedContent); err != nil {
		return true, err
	}
	log.Println("generated:", dstFilePath)
	return true, nil
}

// generateInitializationFile generates `logic.go`.
func (g *genService) generateInitializationFile(importSrcPackages []string) (err error) {
	var (
		logicPackageName = fstr.ToLower(ffile.Basename(g.srcFolder))
		logicFilePath    = ffile.Join(g.srcFolder, logicPackageName+".go")
		logicImports     string
		generatedContent string
	)
	if !utils.IsFileDoNotEdit(logicFilePath) {
		log.Printf("ignore file as it is manually maintained: %s\n", logicFilePath)
		return nil
	}
	for _, importSrcPackage := range importSrcPackages {
		logicImports += fmt.Sprintf(`%s_ "%s"%s`, "\t", importSrcPackage, "\n")
	}
	generatedContent = fstr.ReplaceByMap(consts.TemplateGenServiceLogicContent, map[string]string{
		"{PackageName}": logicPackageName,
		"{Imports}":     logicImports,
		"{CreatedAt}":   utils.CreatedAt(),
	})
	if err = ffile.PutContents(logicFilePath, generatedContent); err != nil {
		return err
	}
	log.Println("generated:", logicFilePath)
	utils.GoFmt(logicFilePath)
	return nil
}

// isToGenerateServiceGoFile checks and returns whether the service content dirty.
func (g *genService) isToGenerateServiceGoFile(dstPackageName, filePath string, funcArray []string) bool {
	var (
		fileContent        = ffile.GetContents(filePath)
		generatedFuncArray = funcArray
		contentFuncArray   = make([]string, 0)
	)
	if fileContent == "" {
		return true
	}
	// remove all comments.
	fileContent, _ = fregex.ReplaceString(`(//.*)|((?s)/\*.*?\*/)`, "", fileContent)
	if fileContent == "" {
		return true
	}
	matches, _ := fregex.MatchAllString(`\s+interface\s+{((?:[^{}]*|\{[^{}]*\})*)\}`, fileContent)
	for _, match := range matches {
		contentFuncArray = append(contentFuncArray, fstr.SplitAndTrim(match[1], "\n")...)
	}
	if len(generatedFuncArray) != len(contentFuncArray) {
		log.Printf("dirty, generated length(%d) != content length(%d): %s\n",
			len(generatedFuncArray), len(contentFuncArray), ffile.RealPath(filePath))
		return true
	}
	sort.Strings(generatedFuncArray)
	sort.Strings(contentFuncArray)
	pattern := fmt.Sprintf(`\*{0,1}%s\.`, dstPackageName)
	for i := 0; i < len(generatedFuncArray); i++ {
		funcDefinition, _ := fregex.ReplaceString(pattern, ``, generatedFuncArray[i])
		if funcDefinition != contentFuncArray[i] {
			log.Printf("dirty, %s != %s: %s\n", funcDefinition, contentFuncArray[i], ffile.RealPath(filePath))
			return true
		}
	}
	return false
}
