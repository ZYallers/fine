package gendao

import (
	"fmt"
	"strings"

	"gitlab.sys.hxsapp.net/hxs/fine/os/fcfg"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fregex"
	"gitlab.sys.hxsapp.net/hxs/fine/text/fstr"
	"gitlab.sys.hxsapp.net/hxs/fine/util/fcast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type localType string

const (
	localTypeUndefined   localType = ""
	localTypeString      localType = "string"
	localTypeDate        localType = "date"
	localTypeDatetime    localType = "datetime"
	localTypeInt         localType = "int"
	localTypeUint        localType = "uint"
	localTypeInt64       localType = "int64"
	localTypeUint64      localType = "uint64"
	localTypeIntSlice    localType = "[]int"
	localTypeInt64Slice  localType = "[]int64"
	localTypeUint64Slice localType = "[]uint64"
	localTypeInt64Bytes  localType = "int64-bytes"
	localTypeUint64Bytes localType = "uint64-bytes"
	localTypeFloat32     localType = "float32"
	localTypeFloat64     localType = "float64"
	localTypeBytes       localType = "[]byte"
	localTypeBool        localType = "bool"
	localTypeJson        localType = "json"
	localTypeJsonb       localType = "jsonb"
)

const (
	fieldTypeBinary     = "binary"
	fieldTypeVarbinary  = "varbinary"
	fieldTypeBlob       = "blob"
	fieldTypeTinyblob   = "tinyblob"
	fieldTypeMediumblob = "mediumblob"
	fieldTypeLongblob   = "longblob"
	fieldTypeInt        = "int"
	fieldTypeTinyint    = "tinyint"
	fieldTypeSmallInt   = "small_int"
	fieldTypeSmallint   = "smallint"
	fieldTypeMediumInt  = "medium_int"
	fieldTypeMediumint  = "mediumint"
	fieldTypeSerial     = "serial"
	fieldTypeBigInt     = "big_int"
	fieldTypeBigint     = "bigint"
	fieldTypeBigserial  = "bigserial"
	fieldTypeReal       = "real"
	fieldTypeFloat      = "float"
	fieldTypeDouble     = "double"
	fieldTypeDecimal    = "decimal"
	fieldTypeMoney      = "money"
	fieldTypeNumeric    = "numeric"
	fieldTypeSmallmoney = "smallmoney"
	fieldTypeBool       = "bool"
	fieldTypeBit        = "bit"
	fieldTypeDate       = "date"
	fieldTypeDatetime   = "datetime"
	fieldTypeTimestamp  = "timestamp"
	fieldTypeTimestampz = "timestamptz"
	fieldTypeJson       = "json"
	fieldTypeJsonb      = "jsonb"
)

// tableField is the struct for table field.
type tableField struct {
	Index   int         // For ordering purpose as map is unordered.
	Name    string      // Field name.
	Type    string      // Field type. Eg: 'int(10) unsigned', 'varchar(64)'.
	Null    bool        // Field can be null or not.
	Key     string      // The index information(empty if it's not an index). Eg: PRI, MUL.
	Default interface{} // Default value for the field.
	Extra   string      // Extra information. Eg: auto_increment.
	Comment string      // Field comment.
}

func (g *genDao) getDB() (*gorm.DB, error) {
	prefix := "fine.gendao.mysql." + g.group + "."
	user := fcfg.GetEnvString(prefix + "user")
	pwd := fcfg.GetEnvString(prefix + "password")
	host := fcfg.GetEnvString(prefix + "host")
	port := fcfg.GetEnvString(prefix + "port")
	database := fcfg.GetEnvString(prefix + "database")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s&maxAllowedPacket=%s&timeout=%s",
		user, pwd, host, port, database, "utf8mb4", "true", "Local", "0", "15s")

	dialect := mysql.Open(dns)
	gormConfig := &gorm.Config{DisableAutomaticPing: true}
	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, fmt.Errorf("db.%s.open.error: %s", g.group, err)
	}
	sdb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("db.%s.db.error: %s", g.group, err)
	}
	if err := sdb.Ping(); err != nil {
		return nil, fmt.Errorf("db.%s.ping.error: %s", g.group, err)
	}
	return db, nil
}

func (g *genDao) getTableFields(db *gorm.DB, tableName string) (map[string]*tableField, error) {
	dest := make([]map[string]interface{}, 0)
	err := db.Raw(fmt.Sprintf("SHOW FULL COLUMNS FROM `%s`", tableName)).Scan(&dest).Error
	if err != nil {
		return nil, err
	}

	fields := make(map[string]*tableField)
	for i, m := range dest {
		field := fcast.ToString(m["Field"])
		fields[field] = &tableField{
			Index:   i,
			Name:    field,
			Type:    fcast.ToString(m["Type"]),
			Null:    fstr.ToUpper(fcast.ToString(m["Null"])) != "NO",
			Key:     fcast.ToString(m["Key"]),
			Default: m["Default"],
			Extra:   fcast.ToString(m["Extra"]),
			Comment: fcast.ToString(m["Comment"]),
		}
	}

	return fields, nil
}

// CheckLocalTypeForField checks and returns corresponding type for given db type.
func (g *genDao) checkLocalTypeForField(fieldType string, fieldValue interface{}) (localType, error) {
	var (
		typeName    string
		typePattern string
	)
	match, _ := fregex.MatchString(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = fstr.Trim(match[1])
		typePattern = fstr.Trim(match[2])
	} else {
		typeName = fstr.Split(fieldType, " ")[0]
	}

	typeName = strings.ToLower(typeName)
	switch typeName {
	case
		fieldTypeBinary,
		fieldTypeVarbinary,
		fieldTypeBlob,
		fieldTypeTinyblob,
		fieldTypeMediumblob,
		fieldTypeLongblob:
		return localTypeBytes, nil
	case
		fieldTypeInt,
		fieldTypeTinyint,
		fieldTypeSmallInt,
		fieldTypeSmallint,
		fieldTypeMediumInt,
		fieldTypeMediumint,
		fieldTypeSerial:
		if fstr.ContainsI(fieldType, "unsigned") {
			return localTypeUint, nil
		}
		return localTypeInt, nil
	case
		fieldTypeBigInt,
		fieldTypeBigint,
		fieldTypeBigserial:
		if fstr.ContainsI(fieldType, "unsigned") {
			return localTypeUint64, nil
		}
		return localTypeInt64, nil
	case
		fieldTypeReal:
		return localTypeFloat32, nil
	case
		fieldTypeDecimal,
		fieldTypeMoney,
		fieldTypeNumeric,
		fieldTypeSmallmoney:
		return localTypeString, nil
	case
		fieldTypeFloat,
		fieldTypeDouble:
		return localTypeFloat64, nil
	case
		fieldTypeBit:
		// It is suggested using bit(1) as boolean.
		if typePattern == "1" {
			return localTypeBool, nil
		}
		s := fcast.ToString(fieldValue)
		// mssql is true|false string.
		if strings.EqualFold(s, "true") || strings.EqualFold(s, "false") {
			return localTypeBool, nil
		}
		if fstr.ContainsI(fieldType, "unsigned") {
			return localTypeUint64Bytes, nil
		}
		return localTypeInt64Bytes, nil
	case
		fieldTypeBool:
		return localTypeBool, nil
	case
		fieldTypeDate:
		return localTypeDate, nil
	case
		fieldTypeDatetime,
		fieldTypeTimestamp,
		fieldTypeTimestampz:
		return localTypeDatetime, nil
	case
		fieldTypeJson:
		return localTypeJson, nil
	case
		fieldTypeJsonb:
		return localTypeJsonb, nil
	default:
		// Auto-detect field type, using key match.
		switch {
		case strings.Contains(typeName, "text") || strings.Contains(typeName, "char") || strings.Contains(typeName, "character"):
			return localTypeString, nil
		case strings.Contains(typeName, "float") || strings.Contains(typeName, "double") || strings.Contains(typeName, "numeric"):
			return localTypeFloat64, nil
		case strings.Contains(typeName, "bool"):
			return localTypeBool, nil
		case strings.Contains(typeName, "binary") || strings.Contains(typeName, "blob"):
			return localTypeBytes, nil
		case strings.Contains(typeName, "int"):
			if fstr.ContainsI(fieldType, "unsigned") {
				return localTypeUint, nil
			}
			return localTypeInt, nil
		case strings.Contains(typeName, "time"):
			return localTypeDatetime, nil
		case strings.Contains(typeName, "date"):
			return localTypeDatetime, nil
		default:
			return localTypeString, nil
		}
	}
}
