package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"hub/util"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hub CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hub CLI %s %s\n", util.CliVersion, runtime.Version())
	},
}
