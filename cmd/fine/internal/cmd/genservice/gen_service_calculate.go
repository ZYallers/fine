package genservice

import (
	"fmt"
	"go/parser"
	"go/token"

	"gitlab.sys.hxsapp.net/hxs/fine/text/fregex"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
)

func (g *genService) calculateCodeCommented(fileContent string, srcCodeCommentedMap map[string]string) error {
	matches, err := fregex.MatchAllString(`((((//.*)|(/\*[\s\S]*?\*/))\s)+)func \((.+?)\) ([\s\S]+?) {`, fileContent)
	if err != nil {
		return err
	}
	for _, match := range matches {
		var (
			structName    string
			structMatch   []string
			funcReceiver  = fstr.Trim(match[1+5])
			receiverArray = fstr.SplitAndTrim(funcReceiver, " ")
			functionHead  = fstr.Trim(fstr.Replace(match[2+5], "\n", ""))
			commentedInfo = ""
		)
		if len(receiverArray) > 1 {
			structName = receiverArray[1]
		} else if len(receiverArray) == 1 {
			structName = receiverArray[0]
		}
		structName = fstr.Trim(structName, "*")
		functionHead = fstr.Replace(functionHead, `,)`, `)`)
		functionHead, _ = fregex.ReplaceString(`\(\s+`, `(`, functionHead)
		functionHead, _ = fregex.ReplaceString(`\s{2,}`, ` `, functionHead)
		if !fstr.IsLetterUpper(functionHead[0]) {
			continue
		}
		// Match and pick the struct name from receiver.
		if structMatch, err = fregex.MatchString("^s([A-Z]\\w+)$", structName); err != nil {
			return err
		}
		if len(structMatch) < 1 {
			continue
		}
		structName = fstr.CaseCamel(structMatch[1])
		commentedInfo = match[1]
		if len(commentedInfo) > 0 {
			srcCodeCommentedMap[fmt.Sprintf("%s-%s", structName, functionHead)] = commentedInfo
		}
	}
	return nil
}

func (g *genService) calculateImportedPackages(fileContent string) (packages []packageItem, err error) {
	f, err := parser.ParseFile(token.NewFileSet(), "", fileContent, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}
	packages = make([]packageItem, 0)
	for _, s := range f.Imports {
		if s.Path != nil {
			if s.Name != nil {
				// If it has alias, and it is not `_`.
				if pkgAlias := s.Name.String(); pkgAlias != "_" {
					packages = append(packages, packageItem{
						Alias:     pkgAlias,
						Path:      s.Path.Value,
						RawImport: pkgAlias + " " + s.Path.Value,
					})
				}
			} else {
				// no alias
				packages = append(packages, packageItem{
					Alias:     "",
					Path:      s.Path.Value,
					RawImport: s.Path.Value,
				})
			}
		}
	}
	return packages, nil
}

func (g *genService) calculateInterfaceFunctions(fileContent string, srcPkgInterfaceMap map[string][]string) (err error) {
	var matches [][]string
	// calculate struct name and its functions according function definitions.
	matches, err = fregex.MatchAllString(`func \((.+?)\) ([\s\S]+?) {`, fileContent)
	if err != nil {
		return err
	}
	for _, match := range matches {
		var (
			structName    string
			structMatch   []string
			funcReceiver  = fstr.Trim(match[1])
			receiverArray = fstr.SplitAndTrim(funcReceiver, " ")
			functionHead  = fstr.Trim(fstr.Replace(match[2], "\n", ""))
		)
		if len(receiverArray) > 1 {
			structName = receiverArray[1]
		} else if len(receiverArray) == 1 {
			structName = receiverArray[0]
		}
		structName = fstr.Trim(structName, "*")
		functionHead = fstr.Replace(functionHead, `,)`, `)`)
		functionHead, _ = fregex.ReplaceString(`\(\s+`, `(`, functionHead)
		functionHead, _ = fregex.ReplaceString(`\s{2,}`, ` `, functionHead)
		if !fstr.IsLetterUpper(functionHead[0]) {
			continue
		}
		// Match and pick the struct name from receiver.
		if structMatch, err = fregex.MatchString("^s([A-Z]\\w+)$", structName); err != nil {
			return err
		}
		if len(structMatch) < 1 {
			continue
		}
		structName = fstr.CaseCamel(structMatch[1])
		if value, ok := srcPkgInterfaceMap[structName]; !ok {
			srcPkgInterfaceMap[structName] = []string{functionHead}
		} else {
			value = append(value, functionHead)
			srcPkgInterfaceMap[structName] = value
		}
	}
	// calculate struct name according type definitions.
	matches, err = fregex.MatchAllString(`type (.+) struct\s*{`, fileContent)
	if err != nil {
		return err
	}
	for _, match := range matches {
		var (
			structName  string
			structMatch []string
		)
		if structMatch, err = fregex.MatchString("^s([A-Z]\\w+)$", match[1]); err != nil {
			return err
		}
		if len(structMatch) < 1 {
			continue
		}
		structName = fstr.CaseCamel(structMatch[1])
		if _, ok := srcPkgInterfaceMap[structName]; !ok {
			srcPkgInterfaceMap[structName] = []string{}
		}
	}
	return
}
