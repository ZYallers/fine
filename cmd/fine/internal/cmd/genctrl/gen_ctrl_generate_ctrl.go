package genctrl

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ZYallers/fine/cmd/fine/internal/consts"
	"github.com/ZYallers/fine/cmd/fine/internal/util/utils"
	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/text/fstr"
)

type controllerGenerator struct{}

func newControllerGenerator() *controllerGenerator {
	return &controllerGenerator{}
}

func (c *controllerGenerator) Generate(dstModuleFolderPath string, apiModuleApiItems []apiItem) (err error) {
	var (
		doneApiItemSet = make(map[string]struct{})
	)
	for _, item := range apiModuleApiItems {
		if _, ok := doneApiItemSet[item.String()]; ok {
			continue
		}
		// retrieve all api items of the same module.
		var (
			subItems   = c.getSubItemsByModuleAndVersion(apiModuleApiItems, item.Module, item.Version)
			importPath = fstr.Replace(ffile.Dir(item.Import), "\\", "/", -1)
		)
		if err = c.doGenerateCtrlNewByModuleAndVersion(dstModuleFolderPath, item.Module, item.Version, importPath); err != nil {
			return
		}

		for _, subItem := range subItems {
			err = c.doGenerateCtrlItem(dstModuleFolderPath, subItem)
			if err != nil {
				return
			}
			doneApiItemSet[subItem.String()] = struct{}{}
		}
	}
	return
}

func (c *controllerGenerator) getSubItemsByModuleAndVersion(items []apiItem, module, version string) (subItems []apiItem) {
	for _, item := range items {
		if item.Module == module && item.Version == version {
			subItems = append(subItems, item)
		}
	}
	return
}

func (c *controllerGenerator) doGenerateCtrlNewByModuleAndVersion(dstModuleFolderPath, module, version string, importPath string) (err error) {
	var (
		moduleFilePath  = filepath.FromSlash(ffile.Join(dstModuleFolderPath, module+".go"))
		ctrlName        = fmt.Sprintf(`c%s`, fstr.UcFirst(version))
		versionFilePath = filepath.FromSlash(ffile.Join(dstModuleFolderPath, version, version+".go"))
		interfaceName   = fmt.Sprintf(`%s.I%s%s`, module, fstr.CaseCamel(module), fstr.UcFirst(version))
	)
	if !ffile.Exists(versionFilePath) {
		content := fstr.ReplaceByMap(consts.TemplateGenCtrlControllerDefineAndNew, f.MapStrStr{
			"{Package}":       version,
			"{CtrlName}":      ctrlName,
			"{InterfaceName}": interfaceName,
			"{CreatedAt}":     utils.CreatedAt(),
			"{ImportPath}":    fmt.Sprintf(`"%s"`, importPath),
		})
		if err = ffile.PutContents(versionFilePath, fstr.TrimLeft(content)); err != nil {
			return err
		}
		log.Println("generated:", moduleFilePath)
	}
	return
}

func (c *controllerGenerator) doGenerateCtrlItem(dstModuleFolderPath string, item apiItem) (err error) {
	var (
		content         string
		methodNameSnake = fstr.CaseConvert(item.MethodName, fstr.SnakeFirstUpper)
		methodFilePath  = filepath.FromSlash(ffile.Join(dstModuleFolderPath, item.Version, fmt.Sprintf(`%s_%s.go`, item.Version, methodNameSnake)))
	)
	if !ffile.Exists(methodFilePath) {
		content = fstr.ReplaceByMap(consts.TemplateGenCtrlControllerMethodFunc, f.MapStrStr{
			"{Module}":          item.Module,
			"{ImportPath}":      item.Import,
			"{CtrlName}":        fmt.Sprintf(`c%s`, fstr.UcFirst(item.Version)),
			"{Version}":         item.Version,
			"{MethodName}":      item.MethodName,
			"{MethodNameSnake}": fstr.CaseConvert(item.MethodName, fstr.SnakeFirstUpper),
			"{CreatedAt}":       utils.CreatedAt(),
		})
		if err = ffile.PutContents(methodFilePath, fstr.TrimLeft(content)); err != nil {
			return err
		}
	}
	log.Println("generated:", methodFilePath)
	return
}
