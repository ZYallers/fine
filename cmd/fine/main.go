package main

import (
	"fmt"
	"log"

	"github.com/ZYallers/fine/cmd/fine/internal/cmd/genctrl"
	"github.com/ZYallers/fine/cmd/fine/internal/cmd/gendao"
	"github.com/ZYallers/fine/cmd/fine/internal/cmd/genservice"
	"github.com/spf13/cobra"
)

const cliVersion = "1.0.10"

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
