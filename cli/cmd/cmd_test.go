package cmd

import (
	"os"
	"testing"

	"github.com/rockwang465/gitlab-client/test"
)

func TestMain(m *testing.M) {
	test.BuildCli()

	os.Exit(m.Run())
}
