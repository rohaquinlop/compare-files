package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Line struct {
	LineNum int
	LineStr string
	Color   int
}

var rootCmd = &cobra.Command{
	Use:   "compare-files",
	Short: "Compare two files and output the differences",
	Long:  `Compare two files and output the differences.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("requires two file arguments")
		}
		return nil
	},
	Run: compareFiles,
}

func BuildDpMemoized(a []string, b []string, dp [][]int) {
	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
}

func reverseLinesSlice(slice []Line) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}

func FindDiffs(file1 []string, file2 []string) []Line {
	res := []Line{}
	n, m := len(file1), len(file2)

	// Init the 2D array with the dimensions of the files n x m
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)

		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	BuildDpMemoized(file1, file2, dp)

	i, j := len(file1), len(file2)

	for i != 0 || j != 0 {
		if i == 0 {
			// Addition
			res = append(res, Line{LineNum: j, LineStr: file2[j-1], Color: 3})
		} else if j == 0 {
			// Delete
			res = append(res, Line{LineNum: i, LineStr: file1[i-1], Color: 2})
		} else if file1[i-1] == file2[j-1] {
			// Common
			res = append(res, Line{LineNum: i, LineStr: file1[i-1], Color: 1})
			i--
			j--
		} else if dp[i-1][j] <= dp[i][j-1] {
			// Addition
			res = append(res, Line{LineNum: j, LineStr: file2[j-1], Color: 3})
			j--
		} else {
			// Delete
			res = append(res, Line{LineNum: i, LineStr: file1[i-1], Color: 2})
			i--
		}
	}

	reverseLinesSlice(res)

	return res
}

func compareFiles(cmd *cobra.Command, args []string) {
	file1_bytes, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file2_bytes, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file1_str := string(file1_bytes)
	file2_str := string(file2_bytes)

	file1_split_str := strings.Split(file1_str, "\n")
	file2_split_str := strings.Split(file2_str, "\n")

	lines := FindDiffs(file1_split_str, file2_split_str)

	deletedColor := color.New(color.BgRed, color.FgWhite)
	addedColor := color.New(color.BgGreen, color.FgWhite)

	for _, line := range lines {
		if line.Color == 1 {
			fmt.Println(line.LineStr)
		} else {
			if line.Color == 2 {
				if strings.TrimSpace(line.LineStr) == "" {
					deletedColor.Printf("- \\n")
				} else {
					deletedColor.Printf("- %s", line.LineStr)
				}
			} else {
				if strings.TrimSpace(line.LineStr) == "" {
					addedColor.Printf("+ \\n")
				} else {
					addedColor.Printf("+ %s", line.LineStr)
				}
			}
			fmt.Println()
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
