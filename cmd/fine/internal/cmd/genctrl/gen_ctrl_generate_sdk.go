package genctrl

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ZYallers/fine/cmd/fine/internal/consts"
	"github.com/ZYallers/fine/cmd/fine/internal/util/utils"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/text/fregex"
	"github.com/ZYallers/fine/text/fstr"
)

type apiSdkGenerator struct{}

func newApiSdkGenerator() *apiSdkGenerator {
	return &apiSdkGenerator{}
}

func (c *apiSdkGenerator) Generate(apiModuleApiItems []apiItem, sdkFolderPath string) (err error) {
	if err = c.doGenerateSdkPkgFile(sdkFolderPath); err != nil {
		return
	}

	var doneApiItemSet = make(map[string]struct{})
	for _, item := range apiModuleApiItems {
		if _, ok := doneApiItemSet[item.String()]; ok {
			continue
		}
		// retrieve all api items of the same module.
		subItems := c.getSubItemsByModuleAndVersion(apiModuleApiItems, item.Module, item.Version)
		if err = c.doGenerateSdkIClient(sdkFolderPath, item.Import, item.Module, item.Version); err != nil {
			return
		}
		if err = c.doGenerateSdkImplementer(subItems, sdkFolderPath, item.Import, item.Module, item.Version); err != nil {
			return
		}
		for _, subItem := range subItems {
			doneApiItemSet[subItem.String()] = struct{}{}
		}
	}
	return
}

func (c *apiSdkGenerator) doGenerateSdkImplementer(items []apiItem, sdkFolderPath, versionImportPath, module, version string) (err error) {
	var (
		pkgName             = ffile.Basename(sdkFolderPath)
		moduleNameCamel     = fstr.CaseCamel(module)
		moduleNameSnake     = fstr.CaseSnake(module)
		moduleImportPath    = fstr.Replace(ffile.Dir(versionImportPath), "\\", "/", -1)
		versionPrefix       = ""
		implementerName     = moduleNameCamel + fstr.UcFirst(version)
		implementerFilePath = filepath.FromSlash(ffile.Join(sdkFolderPath, fmt.Sprintf(
			`%s_%s_%s.go`, pkgName, moduleNameSnake, version,
		)))
	)
	// implementer file template.
	var importPaths = make([]string, 0)
	importPaths = append(importPaths, fmt.Sprintf("\t\"%s\"", moduleImportPath))
	importPaths = append(importPaths, fmt.Sprintf("\t\"%s\"", versionImportPath))
	implementerFileContent := fstr.TrimLeft(fstr.ReplaceByMap(consts.TemplateGenCtrlSdkImplementer, f.MapStrStr{
		"{PkgName}":         pkgName,
		"{ImportPaths}":     fstr.Join(importPaths, "\n"),
		"{ImplementerName}": implementerName,
		"{CreatedAt}":       utils.CreatedAt(),
	}))
	implementerFileContent += fstr.TrimLeft(fstr.ReplaceByMap(consts.TemplateGenCtrlSdkImplementerNew, f.MapStrStr{
		"{Module}":          module,
		"{VersionPrefix}":   versionPrefix,
		"{ImplementerName}": implementerName,
	}))
	// implementer functions definitions.
	for _, item := range items {
		implementerFileContent += fstr.TrimLeft(fstr.ReplaceByMap(consts.TemplateGenCtrlSdkImplementerFunc, f.MapStrStr{
			"{Version}":         item.Version,
			"{MethodName}":      item.MethodName,
			"{ImplementerName}": implementerName,
		}))
		implementerFileContent += "\n"
	}
	err = ffile.PutContents(implementerFilePath, implementerFileContent)
	log.Println("generated:", implementerFilePath)
	return
}

func (c *apiSdkGenerator) getSubItemsByModuleAndVersion(items []apiItem, module, version string) (subItems []apiItem) {
	for _, item := range items {
		if item.Module == module && item.Version == version {
			subItems = append(subItems, item)
		}
	}
	return
}

func (c *apiSdkGenerator) doGenerateSdkIClient(sdkFolderPath, versionImportPath, module, version string) (err error) {
	var (
		fileContent             string
		isDirty                 bool
		isExist                 bool
		pkgName                 = ffile.Basename(sdkFolderPath)
		funcName                = fstr.CaseCamel(module) + fstr.UcFirst(version)
		interfaceName           = fmt.Sprintf(`I%s`, funcName)
		moduleImportPath        = fstr.Replace(fmt.Sprintf(`"%s"`, ffile.Dir(versionImportPath)), "\\", "/", -1)
		iClientFilePath         = filepath.FromSlash(ffile.Join(sdkFolderPath, fmt.Sprintf(`%s.iclient.go`, pkgName)))
		interfaceFuncDefinition = fmt.Sprintf(
			`%s() %s.%s`,
			fstr.CaseCamel(module)+fstr.UcFirst(version), module, interfaceName,
		)
	)
	if isExist = ffile.Exists(iClientFilePath); isExist {
		fileContent = ffile.GetContents(iClientFilePath)
	} else {
		fileContent = fstr.TrimLeft(fstr.ReplaceByMap(consts.TemplateGenCtrlSdkIClient, f.MapStrStr{
			"{PkgName}":   pkgName,
			"{CreatedAt}": utils.CreatedAt(),
		}))
	}

	// append the import path to current import paths.
	if !fstr.Contains(fileContent, moduleImportPath) {
		isDirty = true
		// It is without using AST, because it is from a template.
		fileContent, err = fregex.ReplaceString(
			`(import \([\s\S]*?)\)`,
			fmt.Sprintf("$1\t%s\n)", moduleImportPath),
			fileContent,
		)
		if err != nil {
			return
		}
	}

	// append the function definition to interface definition.
	if !fstr.Contains(fileContent, interfaceFuncDefinition) {
		isDirty = true
		// It is without using AST, because it is from a template.
		fileContent, err = fregex.ReplaceString(
			`(type IClient interface {[\s\S]*?)}`,
			fmt.Sprintf("$1\t%s\n}", interfaceFuncDefinition),
			fileContent,
		)
		if err != nil {
			return
		}
	}
	if isDirty {
		err = ffile.PutContents(iClientFilePath, fileContent)
		if isExist {
			log.Println("updated:", iClientFilePath)
		} else {
			log.Println("generated:", iClientFilePath)
		}
	}
	return
}

func (c *apiSdkGenerator) doGenerateSdkPkgFile(sdkFolderPath string) (err error) {
	var (
		pkgName     = ffile.Basename(sdkFolderPath)
		pkgFilePath = filepath.FromSlash(ffile.Join(sdkFolderPath, fmt.Sprintf(`%s.go`, pkgName)))
		fileContent string
	)
	if ffile.Exists(pkgFilePath) {
		return nil
	}
	fileContent = fstr.TrimLeft(fstr.ReplaceByMap(consts.TemplateGenCtrlSdkPkgNew, f.MapStrStr{
		"{PkgName}":   pkgName,
		"{CreatedAt}": utils.CreatedAt(),
	}))
	err = ffile.PutContents(pkgFilePath, fileContent)
	log.Println("generated:", pkgFilePath)
	return
}
