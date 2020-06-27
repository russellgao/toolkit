package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgfileName string
	showVersion bool
)

var RootCmd = &cobra.Command{
	Use:   "tkctl",
	Short: "tkctl is a toolkit entrypoint,run `tkctl --help` get more information.",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			printVersion()
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(NewTextReplaceCmd())
	RootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "show the version and exit")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
