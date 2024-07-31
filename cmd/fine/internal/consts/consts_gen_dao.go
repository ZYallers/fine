package consts

const TemplateGenDaoIndexContent = `
// =================================================================================
// This is auto-generated by Fine CLI tool only once. Fill this file as you wish.
// {TplCreatedAt}.
// =================================================================================

package {TplPackageName}

import (
	"{TplImportPrefix}/internal"
)

// internal{TplTableNameCamelCase}Dao is internal type for wrapping internal dao implements.
type internal{TplTableNameCamelCase}Dao = *internal.{TplTableNameCamelCase}Dao

// {TplTableNameCamelLowerCase}Dao is the data access object for table {TplTableName}.
// You can define custom methods on it to extend its functionality as you wish.
type {TplTableNameCamelLowerCase}Dao struct {
	internal{TplTableNameCamelCase}Dao
}

var (
	// {TplTableNameCamelCase} is globally public accessible object for table {TplTableName} operations.
	{TplTableNameCamelCase} = {TplTableNameCamelLowerCase}Dao{
		internal.New{TplTableNameCamelCase}Dao(),
	}
)
`

const TemplateGenDaoInternalContent = `
// ==========================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// {TplCreatedAt}.
// ==========================================================================

package internal

import (
	"github.com/ZYallers/fine/database/fmysql"
	"gorm.io/gorm"
)

// {TplTableNameCamelCase}Dao is the data access object for table {TplTableName}.
type {TplTableNameCamelCase}Dao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns {TplTableNameCamelCase}Columns // columns contains all the column names of Table for convenient usage.
}

// {TplTableNameCamelCase}Columns defines and stores column names for table {TplTableName}.
type {TplTableNameCamelCase}Columns struct {
	{TplColumnDefine}
}

// {TplTableNameCamelLowerCase}Columns holds the columns for table {TplTableName}.
var {TplTableNameCamelLowerCase}Columns = {TplTableNameCamelCase}Columns{
	{TplColumnNames}
}

// New{TplTableNameCamelCase}Dao creates and returns a new dao object for table data access.
func New{TplTableNameCamelCase}Dao() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{
		group:   "{TplGroupName}",
		table:   "{TplTableName}",
		columns: {TplTableNameCamelLowerCase}Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current dao.
func (dao *{TplTableNameCamelCase}Dao) DB() *gorm.DB { return fmysql.DB(dao.group) }

// Table returns the table name of current dao.
func (dao *{TplTableNameCamelCase}Dao) Table() string { return dao.table }

// Columns returns all column names of current dao.
func (dao *{TplTableNameCamelCase}Dao) Columns() {TplTableNameCamelCase}Columns { return dao.columns }

// Group returns the configuration group name of database of current dao.
func (dao *{TplTableNameCamelCase}Dao) Group() string { return dao.group }

// Session creates and returns the session for current dao.
func (dao *{TplTableNameCamelCase}Dao) Session() *gorm.DB { return dao.DB().Table(dao.table) }

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not commit or rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *{TplTableNameCamelCase}Dao) Transaction(f func(tx *gorm.DB) error) error { return dao.DB().Transaction(f) }
`

const TemplateGenDaoEntityContent = `
// =================================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// {TplCreatedAt}.
// =================================================================================

package {TplPackageName}

{TplPackageImports}

// {TplTableNameCamelCase} is the golang structure for table {TplTableName}.
{TplStructDefine}
`
