package genctrl

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ZYallers/fine/cmd/fine/internal/consts"
	"github.com/ZYallers/fine/cmd/fine/internal/util/utils"
	"github.com/ZYallers/fine/container/fmap"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/text/fstr"
)

type apiInterfaceGenerator struct{}

func newApiInterfaceGenerator() *apiInterfaceGenerator {
	return &apiInterfaceGenerator{}
}

func (c *apiInterfaceGenerator) Generate(apiModuleFolderPath string, apiModuleApiItems []apiItem) (err error) {
	if len(apiModuleApiItems) == 0 {
		return nil
	}
	var firstApiItem = apiModuleApiItems[0]
	if err = c.doGenerate(apiModuleFolderPath, firstApiItem.Module, apiModuleApiItems); err != nil {
		return
	}
	return
}

func (c *apiInterfaceGenerator) doGenerate(apiModuleFolderPath string, module string, items []apiItem) (err error) {
	var (
		moduleFilePath = filepath.FromSlash(ffile.Join(apiModuleFolderPath, fmt.Sprintf(`%s.go`, module)))
		importPathMap  = fmap.NewStrAnyMap()
		importPaths    []string
	)
	// if there's already exist file that with the same but not auto generated go file,
	// it uses another file name.
	if !utils.IsFileDoNotEdit(moduleFilePath) {
		moduleFilePath = filepath.FromSlash(ffile.Join(apiModuleFolderPath, fmt.Sprintf(`%s.if.go`, module)))
	}
	// all import paths.
	importPathMap.Set("\t"+`"github.com/gin-gonic/gin"`+"\n", struct{}{})
	for _, item := range items {
		importPathMap.Set(fmt.Sprintf("\t"+`"%s"`, item.Import), struct{}{})
	}
	importPaths = importPathMap.Keys()
	// interface definitions.
	var (
		doneApiItemSet      = make(map[string]struct{})
		interfaceDefinition string
		interfaceContent    = fstr.TrimLeft(fstr.ReplaceByMap(consts.TemplateGenCtrlApiInterface, f.MapStrStr{
			"{Module}":      module,
			"{ImportPaths}": fstr.Join(importPaths, "\n"),
			"{CreatedAt}":   utils.CreatedAt(),
		}))
	)
	for _, item := range items {
		if _, ok := doneApiItemSet[item.String()]; ok {
			continue
		}
		// retrieve all api items of the same module.
		subItems := c.getSubItemsByModuleAndVersion(items, item.Module, item.Version)
		var (
			method        string
			methods       = make([]string, 0)
			interfaceName = fmt.Sprintf(`I%s%s`, fstr.CaseCamel(item.Module), fstr.UcFirst(item.Version))
		)
		for _, subItem := range subItems {
			method = fmt.Sprintf(
				"\t%s(ctx *gin.Context, req *%s.%sReq) (res *%s.%sRes, err error)",
				subItem.MethodName, subItem.Version, subItem.MethodName, subItem.Version, subItem.MethodName,
			)
			methods = append(methods, method)
			doneApiItemSet[subItem.String()] = struct{}{}
		}
		interfaceDefinition += fmt.Sprintf("type %s interface {", interfaceName)
		interfaceDefinition += "\n"
		interfaceDefinition += fstr.Join(methods, "\n")
		interfaceDefinition += "\n"
		interfaceDefinition += fmt.Sprintf("}")
		interfaceDefinition += "\n\n"
	}
	interfaceContent = fstr.TrimLeft(fstr.ReplaceByMap(interfaceContent, f.MapStrStr{
		"{Interfaces}": fstr.TrimRightStr(interfaceDefinition, "\n", 2),
	}))
	err = ffile.PutContents(moduleFilePath, interfaceContent)
	log.Println("generated:", moduleFilePath)
	return
}

func (c *apiInterfaceGenerator) getSubItemsByModule(items []apiItem, module string) (subItems []apiItem) {
	for _, item := range items {
		if item.Module == module {
			subItems = append(subItems, item)
		}
	}
	return
}

func (c *apiInterfaceGenerator) getSubItemsByModuleAndVersion(items []apiItem, module, version string) (subItems []apiItem) {
	for _, item := range items {
		if item.Module == module && item.Version == version {
			subItems = append(subItems, item)
		}
	}
	return
}
