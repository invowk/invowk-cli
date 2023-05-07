package _type

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestPromptText_CheckForRequired(t *testing.T) {
	pt := PromptText("")
	errs := pt.Check()

	require.NotEmpty(t, errs)
	require.Error(t, errs[0])
	require.Equal(t, "PromptText is a required field", errs[0].Error())
}

func TestPromptText_CheckForMaxSize(t *testing.T) {
	pt := PromptText(strings.Repeat("A", 81))
	errs := pt.Check()

	require.NotEmpty(t, errs)
	require.Error(t, errs[0])
	require.Equal(t, "PromptText must be a maximum of 80 characters in length", errs[0].Error())
}

func TestHelpText_CheckForRequired(t *testing.T) {
	ht := HelpText("")
	errs := ht.Check()

	require.NotEmpty(t, errs)
	require.Error(t, errs[0])
	require.Equal(t, "HelpText is a required field", errs[0].Error())
}
