package group

import (
	"smoothcloudcli/json"
	"smoothcloudcli/prompt"
)

func CreateGroup() {
	name := prompt.Input("Name of the group", "")
	templateName := prompt.InputWithSelect("Which template do you want to use?", []string{})
	groupConfig := json.GroupConfig{
		Name: name,
		TemplateName: templateName,
	}
}