package cmd

import (
	out "github.com/rockwang465/go-gitlab-client/cli/output"
	"github.com/rockwang465/go-gitlab-client/gitlab"
)

func printMeta(meta *gitlab.ResponseMeta, withPagination bool) {
	if verbose {
		out.Meta(meta, withPagination)
	}
}
