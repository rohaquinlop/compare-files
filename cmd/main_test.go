package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"version"})
	rootCmd.Execute()

	expected := "compare-files v0.1 -- HEAD\n"

	assert.Equal(t, expected, actual.String())
}

type findDiffsTest struct {
	arg1, arg2 []string
	expected   []Line
}

var findDiffsTests = []findDiffsTest{
	{
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"},
		[]Line{
			{1, "a", 1},
			{2, "b", 1},
			{3, "c", 1},
		},
	},
	{
		[]string{"a", "b", "c"},
		[]string{"a", "b", "d"},
		[]Line{
			{1, "a", 1},
			{2, "b", 1},
			{3, "c", 2},
			{3, "d", 3},
		},
	},
	{
		[]string{
			"\"[go]\": {",
			"    \"editor.formatOnSave\": true,",
			"    \"editor.codeActionsOnSave\": {",
			"        \"source.organizeImports\": false",
			"    },",
			"}",
		},
		[]string{
			"\"[go]\": {",
			"    \"editor.formatOnSave\": true,",
			"    \"editor.codeActionsOnSave\": {",
			"        \"source.organizeImports\": true",
			"    },",
			"}",
			"\n",
		},
		[]Line{
			{1, "\"[go]\": {", 1},
			{2, "    \"editor.formatOnSave\": true,", 1},
			{3, "    \"editor.codeActionsOnSave\": {", 1},
			{4, "        \"source.organizeImports\": false", 2},
			{4, "        \"source.organizeImports\": true", 3},
			{5, "    },", 1},
			{6, "}", 1},
			{7, "\n", 3},
		},
	},
}

func TestFindDiffs(t *testing.T) {
	for _, test := range findDiffsTests {
		actual := FindDiffs(test.arg1, test.arg2)
		assert.Equal(t, test.expected, actual)
	}
}
