package genctrl

import (
	"fmt"

	"github.com/ZYallers/fine/os/ffile"
)

type genCtrl struct {
	srcFolder         string
	dstFolder         string
	sdkPath           string
	projectModuleName string
}

func (g *genCtrl) Run() error {
	if !ffile.Exists(g.srcFolder) {
		return fmt.Errorf("source folder path \"%s\" not exist", g.srcFolder)
	}
	if !ffile.Exists(g.dstFolder) {
		return fmt.Errorf("destination folder path \"%s\" not exist", g.dstFolder)
	}
	// retrieve all api modules.
	apiModuleFolderPaths, err := ffile.ScanDir(g.srcFolder, "*", false)
	if err != nil {
		return err
	}
	for _, apiModuleFolderPath := range apiModuleFolderPaths {
		if !ffile.IsDir(apiModuleFolderPath) {
			continue
		}
		// generate go files by api module.
		var (
			module              = ffile.Basename(apiModuleFolderPath)
			dstModuleFolderPath = ffile.Join(g.dstFolder, module)
		)
		err = g.generateByModule(apiModuleFolderPath, dstModuleFolderPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// generateByModule parses certain api and generate associated go files by certain module, not all api modules.
func (g *genCtrl) generateByModule(apiModuleFolderPath, dstModuleFolderPath string) (err error) {
	// parse src and dst folder go files.
	apiItemsInSrc, err := g.getApiItemsInSrc(apiModuleFolderPath)
	if err != nil {
		return err
	}
	apiItemsInDst, err := g.getApiItemsInDst(dstModuleFolderPath)
	if err != nil {
		return err
	}

	// generate api interface go files.
	if err = newApiInterfaceGenerator().Generate(apiModuleFolderPath, apiItemsInSrc); err != nil {
		return
	}

	// generate controller go files.
	// api filtering for already implemented api controllers.
	var (
		alreadyImplementedCtrlSet = make(map[string]struct{})
		toBeImplementedApiItems   = make([]apiItem, 0)
	)
	for _, item := range apiItemsInDst {
		alreadyImplementedCtrlSet[item.String()] = struct{}{}
	}
	for _, item := range apiItemsInSrc {
		if _, ok := alreadyImplementedCtrlSet[item.String()]; ok {
			continue
		}
		toBeImplementedApiItems = append(toBeImplementedApiItems, item)
	}
	if len(toBeImplementedApiItems) > 0 {
		err = newControllerGenerator().Generate(dstModuleFolderPath, toBeImplementedApiItems)
		if err != nil {
			return
		}
	}

	// generate sdk go files.
	if g.sdkPath != "" {
		if err = newApiSdkGenerator().Generate(apiItemsInSrc, g.sdkPath); err != nil {
			return
		}
	}
	return
}
