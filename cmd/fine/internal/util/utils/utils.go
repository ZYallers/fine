package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"gitlab.sys.hxsapp.net/hxs/fine/os/ffile"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fregex"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
	"golang.org/x/tools/imports"
)

func CreatedAt() string {
	return fmt.Sprintf(`Created at %s`, time.Now().Format("2006/01/02 15:04:05.000"))
}

func GetProjectModuleName() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	goModPath := ffile.Join(pwd, "go.mod")
	if !ffile.Exists(goModPath) {
		return "", fmt.Errorf("file \"%s\" not exist", goModPath)
	}
	pattern := `^module\s+(.+)\s*`
	if regex, err := regexp.Compile(pattern); err != nil {
		return "", fmt.Errorf(`regexp compile failed for pattern "%s" error: %v`, pattern, err)
	} else {
		match := regex.FindStringSubmatch(ffile.GetContents(goModPath))
		return fstr.Trim(match[1]), nil
	}
}

// RemoveDuplicateWithString Remove duplicate values
func RemoveDuplicateWithString(arr []string) []string {
	var result []string
	tmp := map[string]byte{}
	for _, val := range arr {
		l := len(tmp)
		tmp[val] = 0
		if len(tmp) != l {
			result = append(result, val)
		}
	}
	return result
}

func GetImportPath(filePath string) string {
	// If `filePath` does not exist, create it firstly to find the import path.
	var realPath = ffile.RealPath(filePath)
	if realPath == "" {
		_ = ffile.Mkdir(filePath)
		realPath = ffile.RealPath(filePath)
	}

	var (
		newDir     = ffile.Dir(realPath)
		oldDir     string
		suffix     string
		goModName  = "go.mod"
		goModPath  string
		importPath string
	)

	if ffile.IsDir(filePath) {
		suffix = ffile.Basename(filePath)
	}
	for {
		goModPath = ffile.Join(newDir, goModName)
		if ffile.Exists(goModPath) {
			match, _ := fregex.MatchString(`^module\s+(.+)\s*`, ffile.GetContents(goModPath))
			importPath = fstr.Trim(match[1]) + "/" + suffix
			importPath = fstr.Replace(importPath, `\`, `/`)
			importPath = fstr.TrimRight(importPath, `/`)
			return importPath
		}
		oldDir = newDir
		newDir = ffile.Dir(oldDir)
		if newDir == oldDir {
			return ""
		}
		suffix = ffile.Basename(oldDir) + "/" + suffix
	}
}

// GetModPath retrieves and returns the file path of go.mod for current project.
func GetModPath() string {
	var (
		oldDir, _ = os.Getwd()
		newDir    = oldDir
		goModName = "go.mod"
		goModPath string
	)
	for {
		goModPath = ffile.Join(newDir, goModName)
		if ffile.Exists(goModPath) {
			return goModPath
		}
		newDir = ffile.Dir(oldDir)
		if newDir == oldDir {
			break
		}
		oldDir = newDir
	}
	return ""
}

// GoFmt formats the source file and adds or removes import statements as necessary.
func GoFmt(path string) {
	replaceFunc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			fmt.Printf("format files \"%s\" error: %v\n", path, err)
			return content
		}
		return string(res)
	}

	var err error
	if ffile.IsFile(path) {
		// File format.
		if ffile.ExtName(path) != "go" {
			return
		}
		err = ffile.ReplaceFileFunc(replaceFunc, path)
	} else {
		// Folder format.
		err = ffile.ReplaceDirFunc(replaceFunc, path, "*.go", true)
	}
	if err != nil {
		log.Printf("format files \"%s\" error: %v\n", path, err)
	}
	log.Println("formatted:", path)
}

// IsFileDoNotEdit checks and returns whether file contains `do not edit` key.
func IsFileDoNotEdit(filePath string) bool {
	if !ffile.Exists(filePath) {
		return true
	}
	return fstr.Contains(ffile.GetContents(filePath), `DO NOT EDIT`)
}
