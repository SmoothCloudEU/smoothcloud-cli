package group

import (
	"fmt"
	"os"
	"path/filepath"
	"smoothcloudcli/command/template"
	"smoothcloudcli/json"
	"smoothcloudcli/prompt"
	"strings"
)

func CreateGroup() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error opening userhomedir: %v\n", err)
		return
	}
	smoothcloudConfig := filepath.Join(userDir, ".smoothcloud", "config.json")
	var config json.MainConfig
	err = json.LoadJSON(smoothcloudConfig, &config)
	if err != nil {
		fmt.Printf("Error loading json: %v\n", err)
		return
	}
	var serviceVersionStruct json.ServiceVersion
	err = json.LoadJSONFromURL("https://github.com/SmoothCloudEU/smoothcloud-manifest/raw/refs/heads/main/serviceVersions.json", &serviceVersionStruct)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	groupType := prompt.InputWithSelect("Type of the group", []string{"Proxy", "Lobby", "Server"})
	name := prompt.Input("Name of the group", "")
	templateName := prompt.InputWithSelect("Which template do you want to use?", func() []string {
		templates := template.GetAllTemplates()
		if len(templates) == 0 {
			return []string{"create"}
		}
		templates = append(templates, "create")
		return templates
	}())
	startPriority := prompt.InputInteger("Which start priority should have services of this group?", "")
	static := prompt.InputWithSelect("Should the group start static services?", []string{"yes", "no"})
	maintenance := prompt.InputWithSelect("Should the group start services automatically in maintenance mode?", []string{"yes", "no"})
	permission := prompt.InputWithEmpty("Which permission is needed to join services of this group? (Leave empty for no permission)", "")
	minMemory := prompt.InputInteger("How many memory should use a service of this group minimal? (in MB)", "")
	maxMemory := prompt.InputInteger("How many memory should use a service of this group maximal?", "")
	minOnlineServices := prompt.InputInteger("How many services of this group should be online minimal?", "")
	maxOnlineServices := prompt.InputInteger("How many services of this group should be online maximal?", "")
	maxPlayers := prompt.InputInteger("How many players should can connect a service of this group?", "")
	newServiceProcent := prompt.InputInteger("How many players are needed on this group to start a new service? (in %)", "")
	serverVersions, proxyVersions := ExtractServiceVersions(serviceVersionStruct)
	serviceVersion := prompt.InputWithSelect("Which service version and software do you want to use?", func() []string {
		switch groupType {
		case "Proxy":
			return proxyVersions
		default:
			return serverVersions
		}
	}())
	java := prompt.InputWithSelect("Which java version do you want to use?", []string{"JAVA_23", "JAVA_21", "JAVA_17", "JAVA_11", "JAVA_8"})
	var staticBool bool
	switch static {
	case "yes":
		staticBool = true
	case "no":
		staticBool = false
	}
	var maintenanceBool bool
	switch maintenance {
	case "yes":
		maintenanceBool = true
	case "no":
		maintenanceBool = false
	}
	groupConfig := json.GroupConfig{
		Name: name,
		TemplateName: func() string {
			if templateName == "create" {
				templateFolderPath := filepath.Join(config.WorkingDirectory, "templates", name)
				os.MkdirAll(templateFolderPath, os.ModePerm)
				templatesPath := filepath.Join(config.WorkingDirectory, "storage", "templates.json")
				err = json.AddNestedKeyValue(templatesPath, "templates." + name, json.TemplateConfig{Name: name, ServiceVersion: serviceVersion})
				return name
			}
			return templateName
		}(),
		StartPriority: startPriority,
		Static: staticBool,
		Maintenance: maintenanceBool,
		Permission: func() any {
			if permission == "" {
				return nil
			}
			return permission
		}(),
		MinMemory: minMemory,
		MaxMemory: maxMemory,
		MinOnlineServices: minOnlineServices,
		MaxOnlineServices: maxOnlineServices,
		MaxPlayers: maxPlayers,
		NewServiceProcent: newServiceProcent,
		ServiceVersion: serviceVersion,
		Java: java,
	}
	var groupTypePath string
	switch groupType {
	case "Proxy":
		groupTypePath = "proxies"
	case "Lobby":
		groupTypePath = "lobbies"
	case "Server":
		groupTypePath = "server"
	}
	groupConfigPath := filepath.Join(config.WorkingDirectory, "groups", groupTypePath, name + ".json")
	err = json.SaveJSON(groupConfigPath, groupConfig)
	if err != nil {
		fmt.Printf("Error saving configuration to %s: %v\n", groupConfigPath, err)
		return
	}
}

func ExtractServiceVersions(config json.ServiceVersion) ([]string, []string) {
	var serviceVersions []string
	var proxyVersions []string
	for serverType, serverConfig := range config.SERVER {
		for version := range serverConfig.Versions {
			serviceVersion := fmt.Sprintf("%s_%s", strings.ToUpper(serverType), version)
			serviceVersions = append(serviceVersions, serviceVersion)
		}
	}
	for serverType, serverConfig := range config.PROXY {
		for version := range serverConfig.Versions {
			serviceVersion := fmt.Sprintf("%s_%s", strings.ToUpper(serverType), version)
			proxyVersions = append(proxyVersions, serviceVersion)
		}
	}
	return serviceVersions, proxyVersions
}
