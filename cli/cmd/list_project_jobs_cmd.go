package cmd

import (
	"fmt"

	"github.com/fatih/color"
	out "github.com/rockwang465/gitlab-client/cli/output"
	"github.com/rockwang465/gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

var projectJobsScope string
var projectJobsPrettyOutput bool

func init() {
	listCmd.AddCommand(listProjectJobsCmd)

	listProjectJobsCmd.Flags().StringVarP(&projectJobsScope, "scope", "s", "", "Scope")
	listProjectJobsCmd.Flags().BoolVar(&projectJobsPrettyOutput, "pretty", false, "Use custom output formatting")
}

func fetchProjectJobs(projectId string) {
	color.Yellow("Fetching project's jobs (project id: %s)…", projectId)

	o := &gitlab.JobsOptions{}
	o.Page = page
	o.PerPage = perPage
	if projectJobsScope != "" {
		o.Scope = []string{projectJobsScope}
	}

	loader.Start()
	collection, meta, err := client.ProjectJobs(projectId, o)
	loader.Stop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(collection.Items) == 0 {
		color.Red("No job found for project %s", projectId)
	} else {
		out.Jobs(output, outputFormat, collection, projectJobsPrettyOutput)
	}

	printMeta(meta, true)

	handlePaginatedResult(meta, func() {
		fetchProjectJobs(projectId)
	})
}

var listProjectJobsCmd = &cobra.Command{
	Use:     resourceCmd("project-jobs", "project"),
	Aliases: []string{"pj"},
	Short:   "List project jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "project", args)
		if err != nil {
			return err
		}

		fetchProjectJobs(ids["project_id"])

		return nil
	},
}
