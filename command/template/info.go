package template

import (
	"fmt"
	"os"
	"path/filepath"
	"smoothcloudcli/json"
)

func GetAllTemplates() []string {
	userDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error opening userhomedir: %v\n", err)
		return []string{}
	}
	smoothcloudConfig := filepath.Join(userDir, ".smoothcloud", "config.json")
	var config json.MainConfig
	err = json.LoadJSON(smoothcloudConfig, &config)
	if err != nil {
		fmt.Printf("Error loading json: %v\n", err)
		return []string{}
	}
	templatesPath := filepath.Join(config.WorkingDirectory, "storage", "templates.json")
	var templatesConfig json.TemplatesConfig
	err = json.LoadJSON(templatesPath, &templatesConfig)
	if err != nil {
		fmt.Printf("Error loading json: %v\n", err)
		return []string{}
	}
	var templates []string
	for template := range templatesConfig.Templates {
		templates = append(templates, template)
	}
	return templates
}