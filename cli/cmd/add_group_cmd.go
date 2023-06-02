package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	out "github.com/rockwang465/gitlab-client/cli/output"
	"github.com/rockwang465/gitlab-client/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(addGroupCmd)
}

var addGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Create a new group",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Creating group…")

		group := gitlab.GroupAddPayload{}

		prompt := promptui.Prompt{
			Label: "Name",
		}
		name, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		group.Name = name

		prompt = promptui.Prompt{
			Label: "Path",
		}
		path, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		group.Path = path

		loader.Start()
		createdGroup, meta, err := client.AddGroup(&group)
		loader.Stop()

		out.Group(output, outputFormat, createdGroup)

		printMeta(meta, false)
	},
}
