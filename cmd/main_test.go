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
