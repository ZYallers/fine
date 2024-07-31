package gendao

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	path         string
	daoPath      string
	entityPath   string
	group        string
	tables       string
	prefix       string
	removePrefix string
)

func init() {
	Cmd.Flags().StringVarP(&path, "path", "p", "internal", `directory path for generated files`)
	Cmd.Flags().StringVarP(&daoPath, "daoPath", "d", "dao", `directory path for storing generated dao files under path`)
	Cmd.Flags().StringVarP(&entityPath, "entityPath", "e", "model/entity", `directory path for storing generated entity files under path`)
	Cmd.Flags().StringVarP(&group, "group", "g", "default", `specifying the configuration group name of database for generated ORM instance, it's not necessary and the default value is "default"`)
	Cmd.Flags().StringVarP(&tables, "tables", "t", "", `generate models for given tables, multiple table names separated with ","`)
	Cmd.Flags().StringVarP(&prefix, "prefix", "f", "", `add prefix for all table of specified link/database tables`)
	Cmd.Flags().StringVarP(&removePrefix, "removePrefix", "r", "", `remove specified prefix of the table, multiple prefix separated with ','`)
}

var Cmd = &cobra.Command{
	Use:   "gendao",
	Short: "Generate relevant data models based on table definitions",
	Run: func(cmd *cobra.Command, args []string) {
		genDao := genDao{
			path:         path,
			daoPath:      daoPath,
			entityPath:   entityPath,
			group:        group,
			tables:       tables,
			prefix:       prefix,
			removePrefix: removePrefix,
		}
		if err := genDao.Run(); err != nil {
			log.Fatal(err)
		}
		log.Println("generate dao finished")
	},
}
