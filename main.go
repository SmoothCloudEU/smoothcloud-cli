package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCommand = &cobra.Command{
		Use:     "smoothcloud",
		Aliases: []string{"sc"},
		Short:   "",
	}
	var infoCommand = &cobra.Command{
		Use:   "info",
		Short: "Gets information about the cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("CLI-INFOS")
		},
	}

	rootCommand.AddCommand(infoCommand)

	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
