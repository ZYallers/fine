package gendao

import (
	"fmt"
	"os"

	"github.com/ZYallers/fine/os/ffile"
	"github.com/ZYallers/fine/text/fstr"
	"github.com/spf13/viper"
)

const (
	hackConfigFile = "hack/config.yaml"
)

type genDao struct {
	path         string
	daoPath      string
	entityPath   string
	group        string
	tables       string
	prefix       string
	removePrefix string
}

func (g *genDao) Run() (err error) {
	if fstr.Trim(g.tables, ",") == "" {
		return fmt.Errorf("tables is empty")
	}

	tableNames := fstr.Split(g.tables, ",")
	newTableNames := make([]string, len(tableNames))
	removePrefixArray := fstr.Split(g.removePrefix, ",")
	for i, tableName := range tableNames {
		newTableName := tableName
		for _, v := range removePrefixArray {
			newTableName = fstr.TrimLeftStr(newTableName, v, 1)
		}
		newTableName = g.prefix + newTableName
		newTableNames[i] = newTableName
	}

	// read config
	if err := g.readConfig(); err != nil {
		return err
	}

	// connect database
	db, err := g.getDB()
	if err != nil {
		return err
	}

	// generate dao
	g.generateDao(db, tableNames, newTableNames)

	// generate entity
	g.generateEntity(db, tableNames, newTableNames)

	return
}

func (g *genDao) readConfig() error {
	pwd, _ := os.Getwd()
	realPath := ffile.RealPath(pwd)
	configFile := ffile.Join(realPath, hackConfigFile)
	if !ffile.Exists(configFile) {
		return fmt.Errorf("config file \"%s\" not exists", configFile)
	}
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config file \"%s\" error: %s\n", configFile, err)
	}
	return nil
}
