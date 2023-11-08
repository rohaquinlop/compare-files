package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compare-files",
	Short: "Compare two files and output the differences",
	Long:  `Compare two files and output the differences.`,
	Run:   compare_files,
}

func compare_files(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Please provide two files to compare")
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
