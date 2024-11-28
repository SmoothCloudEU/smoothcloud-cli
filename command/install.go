package command

import (
	"fmt"
	"os"
	"path/filepath"
	"smoothcloudcli/json"
	"strconv"

	"github.com/manifoldco/promptui"
)

const (
	DirGroups    = "/groups"
	DirProxies   = "/groups/proxies"
	DirLobbies   = "/groups/lobbies"
	DirServers   = "/groups/servers"
	DirStorage   = "/storage"
	DirTemplates = "/templates"
)

type Config struct {
	Language string `json:"language"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Memory   int    `json:"memory"`
}

func Install() {
	fmt.Printf("Installing cloud...\n\n")
	existingDirs := getDirectoriesFromCurrentPath()
	selectedDir := SelectDirectory("Select a directory for the cloud installation:", existingDirs)
	if selectedDir == "" {
		fmt.Println("Installation aborted.")
		return
	}
	err := os.MkdirAll(selectedDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", selectedDir, err)
		return
	}
	directories := []string{DirGroups, DirProxies, DirLobbies, DirServers, DirStorage, DirTemplates}
	for _, value := range directories {
		dir := selectedDir + value
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}
	language := InputWithSelect("Which language do you want to use for the cloud?", []string{"en_US", "de_DE"})
	host := Input("Which IPv4 address should the cloud use?", "10.0.0.11")
	port := promptPort()
	memory := promptMemory()
	config := Config{
		Language: language,
		Host:     host,
		Port:     port,
		Memory:   memory,
	}
	configPath := selectedDir + "/config.json"
	err = json.SaveJSON(configPath, config)
	if err != nil {
		fmt.Printf("Error saving configuration to %s: %v\n", configPath, err)
		return
	}
	fmt.Printf("\nCloud has been successfully installed into \"%s\"!\n", selectedDir)
}

func SelectDirectory(question string, directories []string) string {
	prompt := promptui.Select{
		Label: question,
		Items: directories,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error selecting directory: %v\n", err)
		return ""
	}
	return result
}

func getDirectoriesFromCurrentPath() []string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return nil
	}
	var dirs []string
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != currentDir {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error reading directories:", err)
		return nil
	}
	if len(dirs) == 0 {
		fmt.Println("No directories found.")
		return nil
	}
	return dirs
}

func Input(question string, defaultValue string) string {
	prompt := promptui.Prompt{
		Label:     question,
		Default:   defaultValue,
		AllowEdit: true,
		Validate:  validateNonEmpty,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return ""
	}
	return result
}

func InputWithSelect(question string, options []string) string {
	prompt := promptui.Select{
		Label: question,
		Items: options,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error selecting input: %v\n", err)
		return ""
	}
	return result
}

func promptPort() int {
	for {
		portStr := Input("Which port should the cloud use?", "8080")
		port, err := strconv.Atoi(portStr)
		if err != nil || port <= 0 || port > 65535 {
			fmt.Println("Invalid port number. Please enter a valid port (1-65535).")
			continue
		}
		return port
	}
}

func promptMemory() int {
	defaultMemory := "4096"
	prompt := promptui.Prompt{
		Label:     "What is the maximum amount of storage the cloud should use (in MB)?",
		Default:   defaultMemory,
		AllowEdit: true,
		Validate:  validateMemory,
	}
	memoryStr, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error reading memory input: %v\n", err)
		return 4096
	}
	memory, err := strconv.Atoi(memoryStr)
	if err != nil || memory <= 0 {
		fmt.Println("Invalid memory size. Please enter a valid number.")
		return promptMemory()
	}
	return memory
}

func validateMemory(input string) error {
	if input == "" {
		return fmt.Errorf("memory cannot be empty")
	}
	_, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("invalid memory value. please enter a valid number")
	}
	return nil
}

func validateNonEmpty(input string) error {
	if input == "" {
		return fmt.Errorf("input cannot be empty")
	}
	return nil
}
