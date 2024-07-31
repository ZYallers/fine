package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.sys.hxsapp.net/hxs/fine/cmd/fine/internal/cmd/genctrl"
	"gitlab.sys.hxsapp.net/hxs/fine/cmd/fine/internal/cmd/gendao"
	"gitlab.sys.hxsapp.net/hxs/fine/cmd/fine/internal/cmd/genservice"
)

const cliVersion = "1.0.9"

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(version(), genctrl.Cmd, genservice.Cmd, gendao.Cmd)
}

func version() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information of current binary",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("fine cli version: %s\n", cliVersion)
		},
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("cmd execute error: %v", err)
	}
}
