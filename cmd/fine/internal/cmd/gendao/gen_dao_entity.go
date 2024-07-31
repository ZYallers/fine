package gendao

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"gitlab.sys.hxsapp.net/hxs/fine/cmd/fine/internal/consts"
	"gitlab.sys.hxsapp.net/hxs/fine/cmd/fine/internal/util/utils"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/f"
	"gitlab.sys.hxsapp.net/hxs/fine/os/ffile"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fregex"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
	"gitlab.sys.hxsapp.net/hxs/fine/util/fcast"
	"gorm.io/gorm"
)

// formatComment formats the comment string to fit the golang code without any lines.
func (g *genDao) formatComment(comment string) string {
	comment = fstr.ReplaceByArray(comment, []string{
		"\n", " ",
		"\r", " ",
	})
	comment = fstr.Replace(comment, `\n`, " ")
	comment = fstr.Trim(comment)
	return comment
}

func (g *genDao) getFieldGormTag(field *tableField) string {
	result := []string{"column:" + field.Name, "type:" + field.Type}
	if field.Null {
		result = append(result, "null")
	} else {
		result = append(result, "not null")
	}
	// field index
	switch fstr.ToUpper(field.Key) {
	case "PRI":
		result = append(result, "primaryKey")
	case "UNI":
		result = append(result, "unique")
	case "MUL":
		result = append(result, "index")
	}
	// field default
	if field.Default != nil {
		def := fcast.ToString(field.Default)
		if def != "" {
			result = append(result, "default:"+fstr.ToLower(def))
		} else {
			result = append(result, "default:''")
		}
	}
	// field extra
	switch fstr.ToLower(field.Extra) {
	case "auto_increment":
		result = append(result, "autoIncrement")
	case "on update current_timestamp":
		result = append(result, "autoUpdateTime")
	}
	return fstr.Join(result, ";")
}

// generateStructFieldDefinition generates and returns the attribute definition for specified field.
func (g *genDao) generateStructFieldDefinition(field *tableField, tableName string) (attrLines []string, appendImport string) {
	var (
		err              error
		localTypeName    localType
		localTypeNameStr string
		jsonTag          = fstr.CaseConvert(field.Name, fstr.CaseTypeMatch(string(fstr.Snake)))
	)

	typeMapping := viper.GetStringMap("fine.gendao.mysql." + g.group + ".typeMapping")
	if len(typeMapping) > 0 {
		var tryTypeName string
		tryTypeMatch, _ := fregex.MatchString(`(.+?)\((.+)\)`, field.Type)
		if len(tryTypeMatch) == 3 {
			tryTypeName = fstr.Trim(tryTypeMatch[1])
		} else {
			tryTypeName = fstr.Split(field.Type, " ")[0]
		}
		if tryTypeName != "" {
			if v, ok := typeMapping[strings.ToLower(tryTypeName)]; ok {
				if tm, ok2 := v.(map[string]string); ok2 && len(tm) > 0 {
					localTypeNameStr = tm["type"]
					appendImport = tm["import"]
				}
			}
		}
	}

	if localTypeNameStr == "" {
		localTypeName, err = g.checkLocalTypeForField(field.Type, nil)
		if err != nil {
			log.Fatalf("check local type for field type \"%s\" error: %v", field.Type, err)
		}
		localTypeNameStr = string(localTypeName)
		switch localTypeName {
		case localTypeDate, localTypeDatetime:
			localTypeNameStr = "*time.Time"
		case localTypeInt64Bytes:
			localTypeNameStr = "int64"
		case localTypeUint64Bytes:
			localTypeNameStr = "uint64"
		// Special type handle.
		case localTypeJson, localTypeJsonb:
			localTypeNameStr = "string"
		}
	}

	const tagKey = "`"
	removeFieldPrefixArray := fstr.SplitAndTrim("", ",")
	newFiledName := field.Name
	for _, v := range removeFieldPrefixArray {
		newFiledName = fstr.TrimLeftStr(newFiledName, v, 1)
	}

	fieldMapping := viper.GetStringMap("fine.gendao.mysql." + g.group + ".fieldMapping")
	if len(fieldMapping) > 0 {
		if v, ok := fieldMapping[fmt.Sprintf("%s.%s", tableName, newFiledName)]; ok {
			if tm, ok2 := v.(map[string]string); ok2 && len(tm) > 0 {
				localTypeNameStr = tm["type"]
				appendImport = tm["import"]
			}
		}
	}

	attrLines = []string{
		"    #" + fstr.CaseCamel(strings.ToLower(newFiledName)),
		" #" + localTypeNameStr,
	}
	attrLines = append(attrLines, fmt.Sprintf(` #%sjson:"%s"`, tagKey, jsonTag))
	attrLines = append(attrLines, fmt.Sprintf(` #gorm:"%s"`, g.getFieldGormTag(field)))
	attrLines = append(attrLines, fmt.Sprintf(`%s`, tagKey))
	attrLines = append(attrLines, fmt.Sprintf(` #// %s`, g.formatComment(field.Comment)))

	for k, v := range attrLines {
		attrLines[k] = v
	}

	return attrLines, appendImport
}

func (g *genDao) generateStructDefinition(structName, tableName string, fieldMap map[string]*tableField) (string, []string) {
	var appendImports []string
	buffer := bytes.NewBuffer(nil)
	array := make([][]string, len(fieldMap))
	names := g.sortFieldKeyForDao(fieldMap)
	for index, name := range names {
		var imports string
		field := fieldMap[name]
		array[index], imports = g.generateStructFieldDefinition(field, tableName)
		if imports != "" {
			appendImports = append(appendImports, imports)
		}
	}
	tw := tablewriter.NewWriter(buffer)
	tw.SetBorder(false)
	tw.SetRowLine(false)
	tw.SetAutoWrapText(false)
	tw.SetColumnSeparator("")
	tw.AppendBulk(array)
	tw.Render()
	stContent := buffer.String()
	// Let's do this hack of table writer for indent!
	stContent = fstr.Replace(stContent, "  #", "")
	stContent = fstr.Replace(stContent, "` ", "`")
	stContent = fstr.Replace(stContent, "``", "")
	buffer.Reset()
	buffer.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	buffer.WriteString(stContent)
	buffer.WriteString("}")
	return buffer.String(), appendImports
}

func (g *genDao) getImportPartContent(source string, appendImports []string) string {
	var packageImportsArray = make([]string, 0)

	// Time package recognition.
	if strings.Contains(source, "time.Time") {
		packageImportsArray = append(packageImportsArray, `"time"`)
	}

	if len(appendImports) > 0 {
		for _, appendImport := range appendImports {
			packageImportsArray = append(packageImportsArray, fmt.Sprintf(`"%s"`, appendImport))
		}
	}

	// Generate and write content to golang file.
	packageImportsStr := ""
	if len(packageImportsArray) > 0 {
		packageImportsStr = fmt.Sprintf("import(\n%s\n)", fstr.Join(packageImportsArray, "\n"))
	}

	return packageImportsStr
}

func (g *genDao) generateEntityContent(tableName, tableNameCamelCase, structDefine string, appendImports []string) string {
	entityContent := fstr.ReplaceByMap(consts.TemplateGenDaoEntityContent, f.MapStrStr{
		"{TplTableName}":          tableName,
		"{TplPackageImports}":     g.getImportPartContent(structDefine, appendImports),
		"{TplTableNameCamelCase}": tableNameCamelCase,
		"{TplStructDefine}":       structDefine,
		"{TplPackageName}":        filepath.Base(g.entityPath),
		"{TplCreatedAt}":          utils.CreatedAt(),
	})
	return entityContent
}

func (g *genDao) generateEntity(db *gorm.DB, tableNames, newTableNames []string) {
	var dirPathEntity = ffile.Join(g.path, g.entityPath)
	for i, tableName := range tableNames {
		fieldMap, err := g.getTableFields(db, tableName)
		if err != nil {
			log.Fatalf("fetching tables fields failed for table '%s':\n%v", tableName, err)
		}

		var (
			newTableName                    = newTableNames[i]
			entityFilePath                  = filepath.FromSlash(ffile.Join(dirPathEntity, fstr.CaseSnake(newTableName)+".go"))
			structName                      = fstr.CaseCamel(strings.ToLower(newTableName))
			structDefinition, appendImports = g.generateStructDefinition(
				structName,
				tableName,
				fieldMap,
			)
			entityContent = g.generateEntityContent(
				newTableName,
				fstr.CaseCamel(strings.ToLower(newTableName)),
				structDefinition,
				appendImports,
			)
		)
		err = ffile.PutContents(entityFilePath, strings.TrimSpace(entityContent))
		if err != nil {
			log.Fatalf("writing content to '%s' failed: %v", entityFilePath, err)
		} else {
			log.Println("generated:", entityFilePath)
			utils.GoFmt(entityFilePath)
		}
	}
}
