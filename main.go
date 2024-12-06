package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"smoothcloudcli/command"
	"smoothcloudcli/command/group"
	"smoothcloudcli/json"
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
	type Test struct {
		Object	 interface{}	`json:"object"`
	}

	type Test2 struct {
		String 	 string	`json:"string"`
	}
	var testCommand = &cobra.Command{
		Use: "test",
		Short: "Testing things",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var subTestCommand = &args[0]
			switch *subTestCommand {
			case "jsonsave":
				json.SaveJSON("test.json", Test{Object: Test{Object: nil}})
			case "jsonupdate":
				json.UpdateKeyValue("test.json", "object", Test{Object: make(map[string]string)})
			case "jsonadd":
				json.AddKeyValue("test.json", "string", Test2{String: "Test"})
			case "jsonaddnested":
				json.AddNestedKeyValue("test.json", "test.string", Test2{String: "Nested Test"})
			case "jsonremove":
				json.RemoveKey("test.json", "object")
			}
		},
	}
	rootCommand.AddCommand(createGroupCommand, infoCommand, installCommand, testCommand)
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
