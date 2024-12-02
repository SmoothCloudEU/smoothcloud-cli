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

type SQLiteDatabaseConfig struct {
	Filename string `json:"filename"`
	Prefix 	 string `json:"prefix"`
}

type MariaDBDatabaseConfig struct {
	Host 	 string `json:"host"`
	Port 	 int 	`json:"port"`
	Database string	`json:"database"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Prefix 	 string `json:"prefix"`
}

type MongoDBDatabaseConfig struct {
	Host 	 string `json:"host"`
	Port 	 int 	`json:"port"`
	Database string	`json:"database"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Prefix 	 string `json:"prefix"`
}

func Install() {
	fmt.Println("Installing cloud...")
	fmt.Println(" ")
	selectedDir := browseDirectories(".")
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
	for _, dir := range directories {
		path := filepath.Join(selectedDir, dir)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", path, err)
			return
		}
	}
	databaseType := InputWithSelect("Which database type do you want to use for the cloud?", []string{"SQLITE", "MARIADB", "MONGODB"})
	var databaseConfig interface{}
	switch databaseType {
	case "MARIADB":
		databaseConfig = MariaDBDatabaseConfig{
			Host: Input("Enter the host of your database server", "127.0.0.1"),
			Port: promptDatabasePort("3306"),
			Database: Input("Enter the database which you want to use", "smoothcloud"),
			Username: Input("Enter the username of your database user", ""),
			Password: Input("Enter the password of your database user", ""),
			Prefix: "smoothcloud_",
		}
	case "MONGODB":
		databaseConfig = MongoDBDatabaseConfig{
			Host: Input("Enter the host of your database server", "127.0.0.1"),
			Port: promptDatabasePort("3306"),
			Database: Input("Enter the database which you want to use", "smoothcloud"),
			Username: Input("Enter the username of your database user", ""),
			Password: Input("Enter the password of your database user", ""),
			Prefix: "smoothcloud_",
		}
	default:
		databaseConfig = SQLiteDatabaseConfig{
			Filename: "sqlite.db",
			Prefix: "smoothcloud_",
		}
	}
	language := InputWithSelect("Enter the language for the cloud", []string{"en_US", "de_DE"})
	host := Input("Enter a IPv4 address which the cloud should use", "127.0.0.1")
	port := promptPort()
	memory := promptMemory()
	config := Config{
		Language: language,
		Host:     host,
		Port:     port,
		Memory:   memory,
	}
	configPath := filepath.Join(selectedDir, "config.json")
	err = json.SaveJSON(configPath, config)
	if err != nil {
		fmt.Printf("Error saving configuration to %s: %v\n", configPath, err)
		return
	}
	databaseConfigPath := filepath.Join(selectedDir, "/storage", "database.json")
	err = json.SaveJSON(databaseConfigPath, databaseConfig)
	if err != nil {
		fmt.Printf("Error saving configuration to %s: %v\n", configPath, err)
		return
	}
	fmt.Printf("\nCloud has been successfully installed into \"%s\"!\n", selectedDir)
}

func browseDirectories(baseDir string) string {
	currentDir := baseDir
	for {
		dirs, err := getSubdirectories(currentDir)
		if err != nil {
			fmt.Printf("Error reading directories: %v\n", err)
			return ""
		}
		if currentDir != baseDir {
			dirs = append([]string{".."}, dirs...)
		}
		dirs = append(dirs, "Select this directory")
		dirs = append(dirs, "Create a directory")
		prompt := promptui.Select{
			Label: fmt.Sprintf("Current directory: %s\nNavigate or select:", currentDir),
			Items: dirs,
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Error selecting directory: %v\n", err)
			return ""
		}
		if result == "Select this directory" {
			return currentDir
		}
		if result == "Create a directory" {
			prompt := promptui.Prompt{
				Label: 	   "Name of the new directory",
				Default:   "cloud/",
				AllowEdit: false,
				Validate:  validateNonEmpty,
			}
			result, err := prompt.Run()
			if err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				return ""
			}
			os.MkdirAll(result, os.ModePerm)
			return browseDirectories(result)
		}
		if result == ".." {
			currentDir = filepath.Dir(currentDir)
			continue
		}
		currentDir = filepath.Join(currentDir, result)
	}
}

func getSubdirectories(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}
	return dirs, nil
}

func Input(question string, defaultValue string) string {
	prompt := promptui.Prompt{
		Label:     question,
		Default:   defaultValue,
		AllowEdit: false,
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

func promptDatabasePort(suggestion string) int {
	for {
		portStr := Input("Enter the port of your database server", suggestion)
		port, err := strconv.Atoi(portStr)
		if err != nil || port <= 0 || port > 65535 {
			fmt.Println("Invalid port number. Please enter a valid port (1-65535).")
			continue
		}
		return port
	}
}

func promptPort() int {
	for {
		portStr := Input("Which port should the cloud use?", "6042")
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
		return fmt.Errorf("invalid memory value. Please enter a valid number")
	}
	return nil
}

func validateNonEmpty(input string) error {
	if input == "" {
		return fmt.Errorf("input cannot be empty")
	}
	return nil
}
