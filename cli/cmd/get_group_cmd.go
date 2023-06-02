package cmd

import (
	"github.com/fatih/color"
	out "github.com/rockwang465/gitlab-client/cli/output"
	"github.com/spf13/cobra"
)

var groupWithCustomAttributes bool

func init() {
	getCmd.AddCommand(getGroupCmd)

	getGroupCmd.Flags().BoolVarP(&groupWithCustomAttributes, "with-custom-attributes", "x", false, "Include custom attributes (admins only)")
}

var getGroupCmd = &cobra.Command{
	Use:     resourceCmd("group", "group"),
	Aliases: []string{"g"},
	Short:   "Get all details of a group",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids, err := config.aliasIdsOrArgs(currentAlias, "group", args)
		if err != nil {
			return err
		}

		color.Yellow("Fetching group (id: %s)…", ids["group_id"])

		loader.Start()
		group, meta, err := client.Group(ids["group_id"], groupWithCustomAttributes)
		loader.Stop()
		if err != nil {
			return err
		}

		out.Group(output, outputFormat, group)

		printMeta(meta, false)

		relatedCommands([]*relatedCommand{
			newRelatedCommand(listGroupMergeRequestsCmd, map[string]string{
				"group_id": ids["group_id"],
			}),
			newRelatedCommand(listGroupVariablesCmd, map[string]string{
				"group_id": ids["group_id"],
			}),
		})

		return nil
	},
}
