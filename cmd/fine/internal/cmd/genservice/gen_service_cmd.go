package genservice

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	srcFolder string
	dstFolder string
)

func init() {
	Cmd.Flags().StringVarP(&srcFolder, "srcFolder", "s", "internal/logic", "source folder path to be parsed")
	Cmd.Flags().StringVarP(&dstFolder, "dstFolder", "d", "internal/service", "destination folder path storing automatically generated go files")
}

var Cmd = &cobra.Command{
	Use:   "genservice",
	Short: "Generate service file by parse struct and associated functions from packages",
	Run: func(cmd *cobra.Command, args []string) {
		g := &genService{
			srcFolder: srcFolder,
			dstFolder: dstFolder,
		}
		if err := g.Run(); err != nil {
			log.Fatal(err)
		}
		log.Println("generate service finished")
	},
}
