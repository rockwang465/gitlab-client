package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/rockwang465/go-gitlab-client/cli/output"
	"github.com/rockwang465/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listProjectVariablesCmd)
}

var listProjectVariablesCmd = &cobra.Command{
	Use:     resourceCmd("project-variables", "project"),
	Aliases: []string{"project-vars", "pv"},
	Short:   "Get list of a project's variables",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching project variables (id: %s)…", ids["project_id"])

		o := &gitlab.PaginationOptions{
			Page:    page,
			PerPage: perPage,
		}

		loader.Start()
		collection, meta, err := client.ProjectVariables(ids["project_id"], o)
		loader.Stop()
		if err != nil {
			return err
		}

		fmt.Println("")
		if len(collection.Items) == 0 {
			color.Red("  No variable found for project %s", ids["project_id"])
		} else {
			out.Variables(output, outputFormat, collection)
		}

		printMeta(meta, true)

		return nil
	},
}
