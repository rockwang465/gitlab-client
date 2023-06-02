package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/rockwang465/go-gitlab-client/cli/output"
	"github.com/rockwang465/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listSshKeysCmd)
}

func fetchSshKeys() {
	color.Yellow("Fetching current user ssh keys…")

	o := &gitlab.PaginationOptions{}
	o.Page = page
	o.PerPage = perPage

	loader.Start()
	collection, meta, err := client.CurrentUserSshKeys(o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No ssh key found")
	} else {
		out.SshKeys(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, fetchSshKeys)
}

var listSshKeysCmd = &cobra.Command{
	Use:     "ssh-keys",
	Aliases: []string{"sk"},
	Short:   "List current user ssh keys",
	Run: func(cmd *cobra.Command, args []string) {
		fetchSshKeys()
	},
}
