package cmd

import (
	"testing"

	"github.com/rockwang465/go-gitlab-client/test"
)

func TestAddCmd(t *testing.T) {
	test.RunCommandTestCases(t, "users", []*test.CommandTestCase{
		{
			[]string{"add"},
			nil,
			//configs["default"],
			"add_help",
			false,
			nil,
		},
		{
			[]string{"add", "--help"},
			nil,
			//configs["default"],
			"add_help",
			false,
			nil,
		},
	})
}
