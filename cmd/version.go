package cmd

import (
	"fmt"
	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "tkctl version",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Println(version.Print("tkctl"))
}
