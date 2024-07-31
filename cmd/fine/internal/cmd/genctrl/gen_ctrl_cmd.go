package genctrl

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	srcFolder string
	dstFolder string
	sdkPath   string
)

func init() {
	Cmd.Flags().StringVarP(&srcFolder, "srcFolder", "s", "api", "source folder path to be parsed")
	Cmd.Flags().StringVarP(&dstFolder, "dstFolder", "d", "internal/controller", "destination folder path storing automatically generated go files")
	Cmd.Flags().StringVarP(&sdkPath, "sdkPath", "k", "", "also generate sdk go files for api definitions to specified directory")
}

var Cmd = &cobra.Command{
	Use:   "genctrl",
	Short: "Generate controller or sdk by defining api structs",
	Run: func(cmd *cobra.Command, args []string) {
		genCtrl := &genCtrl{
			srcFolder: srcFolder,
			dstFolder: dstFolder,
			sdkPath:   sdkPath,
		}
		if err := genCtrl.Run(); err != nil {
			log.Fatal(err)
		}
		log.Println("generate controller or sdk finished")
	},
}
