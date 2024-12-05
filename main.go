package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"smoothcloudcli/command"
	"smoothcloudcli/command/group"
)

func main() {
	var rootCommand = &cobra.Command{
		Use:     "smoothcloud",
		Aliases: []string{"sc"},
	}
	var createGroupCommand = &cobra.Command{
		Use:   "creategroup",
		Short: "Creates a group",
		Run: func(cmd *cobra.Command, args []string) {
			group.CreateGroup()
		},
	}
	var infoCommand = &cobra.Command{
		Use:   "info",
		Short: "Gets information about the cli",
		Run: func(cmd *cobra.Command, args []string) {
			command.Info()
		},
	}
	var installCommand = &cobra.Command{
		Use:   "install",
		Short: "Will install the cloud",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			command.Install()
		},
	}
	rootCommand.AddCommand(createGroupCommand, infoCommand, installCommand)
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
