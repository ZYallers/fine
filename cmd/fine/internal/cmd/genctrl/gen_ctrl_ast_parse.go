package genctrl

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	"gitlab.sys.hxsapp.net/hxs/fine/os/ffile"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
)

// getStructsNameInSrc retrieves all struct names that end in "Req" and have "fine.Meta" in their body.
func (g *genCtrl) getStructsNameInSrc(filePath string) (structsName []string, err error) {
	var (
		fileContent = ffile.GetContents(filePath)
		fileSet     = token.NewFileSet()
	)
	node, err := parser.ParseFile(fileSet, "", fileContent, parser.ParseComments)
	if err != nil {
		return
	}
	ast.Inspect(node, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			methodName := typeSpec.Name.Name
			if !fstr.HasSuffix(methodName, "Req") {
				// ignore struct name that do not end in "Req"
				return true
			}
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				var buf bytes.Buffer
				if err := printer.Fprint(&buf, fileSet, structType); err != nil {
					return false
				}
				// ignore struct name that match a request, but has no fine.Meta in its body.
				if !fstr.Contains(buf.String(), "f.Meta") {
					return true
				}
				structsName = append(structsName, methodName)
			}
		}
		return true
	})
	return
}

// getImportsInDst retrieves all import paths in the file.
func (g *genCtrl) getImportsInDst(filePath string) (imports []string, err error) {
	var (
		fileContent = ffile.GetContents(filePath)
		fileSet     = token.NewFileSet()
	)

	node, err := parser.ParseFile(fileSet, "", fileContent, parser.ParseComments)
	if err != nil {
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if imp, ok := n.(*ast.ImportSpec); ok {
			imports = append(imports, imp.Path.Value)
		}
		return true
	})

	return
}
