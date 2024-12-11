package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"smoothcloudcli/command"
)

func main() {
	var rootCommand = &cobra.Command{
		Use:     "smoothcloud",
		Aliases: []string{"sc"},
	}
	var createCommand = &cobra.Command{
		Use:   "create",
	}
	var createGroupCommand = &cobra.Command{
		Use:   "group",
		Short: "Create a group with a configurator",
		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}
	var infoCommand = &cobra.Command{
		Use:   "info",
		Short: "Gets information about the cli",
		Run: func(cmd *cobra.Command, args []string) {
			command.Info()
		},
	}
	var setupCommand = &cobra.Command{
		Use:   "setup",
		Short: "Will install the cloud",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			command.Setup()
		},
	}
	createCommand.AddCommand(createGroupCommand)
	rootCommand.AddCommand(createCommand, infoCommand, setupCommand)
	cobra.OnInitialize(func() {
		cobra.EnableCommandSorting = false
	})
	rootCommand.CompletionOptions.DisableDefaultCmd = true
	rootCommand.SetHelpTemplate(`
{{.UseLine}}

Available Commands:
{{range .Commands}}{{if .IsAvailableCommand}}{{.Name | printf "%-15s"}}{{.Short}}{{end}}
{{range .Commands}}{{if .IsAvailableCommand}}{{printf "→ "}}{{.Name | printf "%-13s"}}{{.Short}}{{end}}{{end}}{{end}}
`)
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
