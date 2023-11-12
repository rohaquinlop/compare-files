package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of compare-files",
	Long:  `All software has versions. This is compare-files's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "compare-files v0.1 -- HEAD\n")
		//fmt.Println("compare-files v0.1 -- HEAD")
	},
}
