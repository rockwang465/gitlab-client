package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/rockwang465/gitlab-client/cli/output"
	"github.com/rockwang465/gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var branchesSearch string

func init() {
	listCmd.AddCommand(listProjectBranchesCmd)

	listProjectBranchesCmd.Flags().StringVarP(&branchesSearch, "search", "s", "", "Search term")
}

func fetchProjectBranches(projectId string) {
	color.Yellow("Fetching project's branches (id: %s)…", projectId)

	o := &gitlab.BranchesOptions{}
	o.Page = page
	o.PerPage = perPage
	if branchesSearch != "" {
		o.Search = branchesSearch
	}

	loader.Start()
	collection, meta, err := client.ProjectBranches(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No branch found for project %s", projectId)
	} else {
		out.Branches(output, outputFormat, collection)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectBranches(projectId)
	})
}

var listProjectBranchesCmd = &cobra.Command{
	Use:     resourceCmd("project-branches", "project"),
	Aliases: []string{"pb"},
	Short:   "List project branches",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectBranches(ids["project_id"])

		return nil
	},
}
