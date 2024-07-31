package gendao

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"gitlab.sys.hxsapp.net/hxs/fine/cmd/fine/internal/consts"
	"gitlab.sys.hxsapp.net/hxs/fine/cmd/fine/internal/util/utils"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/f"
	"gitlab.sys.hxsapp.net/hxs/fine/os/ffile"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
	"gorm.io/gorm"
)

func (g *genDao) generateDao(db *gorm.DB, tableNames, newTableNames []string) {
	var (
		dirPathDao         = ffile.Join(g.path, g.daoPath)
		dirPathDaoInternal = ffile.Join(dirPathDao, "internal")
	)
	for i := 0; i < len(tableNames); i++ {
		g.generateDaoSingle(db, tableNames[i], newTableNames[i], dirPathDao, dirPathDaoInternal)
	}
}

func (g *genDao) generateDaoSingle(db *gorm.DB, tableName, newTableName, dirPathDao, dirPathDaoInternal string) {
	// Generating table data preparing.
	fieldMap, err := g.getTableFields(db, tableName)
	if err != nil {
		log.Printf("get table fields error: %s\n", err)
		return
	}
	var (
		tableNameCamelCase      = fstr.CaseCamel(strings.ToLower(newTableName))
		tableNameCamelLowerCase = fstr.CaseCamelLower(strings.ToLower(newTableName))
		tableNameSnakeCase      = fstr.CaseSnake(newTableName)
		importPrefix            = utils.GetImportPath(ffile.Join(g.path, g.daoPath))
	)

	fileName := fstr.Trim(tableNameSnakeCase, "-_.")
	if len(fileName) > 5 && fileName[len(fileName)-5:] == "_test" {
		// Add suffix to avoid the table name which contains "_test",
		// which would make the go file a testing file.
		fileName += "_table"
	}

	// dao - index
	g.generateDaoIndex(
		tableNameCamelCase,
		tableNameCamelLowerCase,
		importPrefix,
		tableName,
		fileName,
		dirPathDao,
	)

	// dao - internal
	g.generateDaoInternal(
		tableNameCamelCase,
		tableNameCamelLowerCase,
		importPrefix,
		fileName,
		tableName,
		dirPathDaoInternal,
		fieldMap,
	)
}

func (g *genDao) sortFieldKeyForDao(fieldMap map[string]*tableField) []string {
	names := make(map[int]string)
	for _, field := range fieldMap {
		names[field.Index] = field.Name
	}
	var (
		i      = 0
		j      = 0
		result = make([]string, len(names))
	)
	for {
		if len(names) == 0 {
			break
		}
		if val, ok := names[i]; ok {
			result[j] = val
			j++
			delete(names, i)
		}
		i++
	}
	return result
}

// generateColumnDefinitionForDao generates and returns the column names definition for specified table.
func (g *genDao) generateColumnDefinitionForDao(fieldMap map[string]*tableField, removeFieldPrefixArray []string) string {
	var (
		buffer = bytes.NewBuffer(nil)
		array  = make([][]string, len(fieldMap))
		names  = g.sortFieldKeyForDao(fieldMap)
	)

	for index, name := range names {
		var (
			field   = fieldMap[name]
			comment = fstr.Trim(fstr.ReplaceByArray(field.Comment, []string{
				"\n", " ",
				"\r", " ",
			}))
		)
		newFiledName := field.Name
		for _, v := range removeFieldPrefixArray {
			newFiledName = fstr.TrimLeftStr(newFiledName, v, 1)
		}
		array[index] = []string{
			"    #" + fstr.CaseCamel(strings.ToLower(newFiledName)),
			" # " + "string",
			" #" + fmt.Sprintf(`// %s`, comment),
		}
	}
	tw := tablewriter.NewWriter(buffer)
	tw.SetBorder(false)
	tw.SetRowLine(false)
	tw.SetAutoWrapText(false)
	tw.SetColumnSeparator("")
	tw.AppendBulk(array)
	tw.Render()
	defineContent := buffer.String()
	// Let's do this hack of table writer for indent!
	defineContent = fstr.Replace(defineContent, "  #", "")
	buffer.Reset()
	buffer.WriteString(defineContent)
	return buffer.String()
}

// generateColumnNamesForDao generates and returns the column names assignment content of column struct for specified table.
func (g *genDao) generateColumnNamesForDao(fieldMap map[string]*tableField, removeFieldPrefixArray []string) string {
	var (
		buffer = bytes.NewBuffer(nil)
		array  = make([][]string, len(fieldMap))
		names  = g.sortFieldKeyForDao(fieldMap)
	)

	for index, name := range names {
		field := fieldMap[name]

		newFiledName := field.Name
		for _, v := range removeFieldPrefixArray {
			newFiledName = fstr.TrimLeftStr(newFiledName, v, 1)
		}

		array[index] = []string{
			"            #" + fstr.CaseCamel(strings.ToLower(newFiledName)) + ":",
			fmt.Sprintf(` #"%s",`, field.Name),
		}
	}
	tw := tablewriter.NewWriter(buffer)
	tw.SetBorder(false)
	tw.SetRowLine(false)
	tw.SetAutoWrapText(false)
	tw.SetColumnSeparator("")
	tw.AppendBulk(array)
	tw.Render()
	namesContent := buffer.String()
	// Let's do this hack of table writer for indent!
	namesContent = fstr.Replace(namesContent, "  #", "")
	buffer.Reset()
	buffer.WriteString(namesContent)
	return buffer.String()
}

func (g *genDao) generateDaoInternal(tableNameCamelCase, tableNameCamelLowerCase, importPrefix,
	fileName, tableName, dirPathDaoInternal string, fieldMap map[string]*tableField) {
	path := filepath.FromSlash(ffile.Join(dirPathDaoInternal, fileName+".go"))
	removeFieldPrefixArray := fstr.SplitAndTrim("", ",")
	modelContent := fstr.ReplaceByMap(consts.TemplateGenDaoInternalContent, f.MapStrStr{
		"{TplImportPrefix}":            importPrefix,
		"{TplTableName}":               tableName,
		"{TplGroupName}":               g.group,
		"{TplTableNameCamelCase}":      tableNameCamelCase,
		"{TplTableNameCamelLowerCase}": tableNameCamelLowerCase,
		"{TplColumnDefine}":            fstr.Trim(g.generateColumnDefinitionForDao(fieldMap, removeFieldPrefixArray)),
		"{TplColumnNames}":             fstr.Trim(g.generateColumnNamesForDao(fieldMap, removeFieldPrefixArray)),
		"{TplCreatedAt}":               utils.CreatedAt(),
	})
	if err := ffile.PutContents(path, strings.TrimSpace(modelContent)); err != nil {
		log.Fatalf("writing content to \"%s\" failed: %v", path, err)
	} else {
		log.Println("generated:", path)
		utils.GoFmt(path)
	}
}

func (g *genDao) generateDaoIndex(tableNameCamelCase, tableNameCamelLowerCase, importPrefix, tableName, fileName, dirPathDao string) {
	path := filepath.FromSlash(ffile.Join(dirPathDao, fileName+".go"))
	indexContent := fstr.ReplaceByMap(consts.TemplateGenDaoIndexContent, f.MapStrStr{
		"{TplImportPrefix}":            importPrefix,
		"{TplTableName}":               tableName,
		"{TplTableNameCamelCase}":      tableNameCamelCase,
		"{TplTableNameCamelLowerCase}": tableNameCamelLowerCase,
		"{TplPackageName}":             filepath.Base(g.daoPath),
		"{TplCreatedAt}":               utils.CreatedAt(),
	})
	if err := ffile.PutContents(path, strings.TrimSpace(indexContent)); err != nil {
		log.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		log.Println("generated:", path)
		utils.GoFmt(path)
	}
}
