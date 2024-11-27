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
	var infoCommand = &cobra.Command{
		Use:   "info",
		Short: "Gets information about the cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("CLI-INFOS")
		},
	}
	var installCommand = &cobra.Command{
		Use:   "install",
		Short: "Installs the cloud into a directory",
		Run: func(cmd *cobra.Command, args []string) {
			command.Install()
		},
	}

	rootCommand.AddCommand(infoCommand)
	rootCommand.AddCommand(installCommand)

	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
