package command

import (
	"fmt"
	"os"
	"path/filepath"
	"smoothcloudcli/json"
	"smoothcloudcli/prompt"
)

const (
	DirGroups    = "/groups"
	DirProxies   = "/groups/proxies"
	DirLobbies   = "/groups/lobbies"
	DirServers   = "/groups/servers"
	DirStorage   = "/storage"
	DirTemplates = "/templates"
)

func Install() {
	fmt.Println("Installing cloud...")
	fmt.Println(" ")
	selectedDir := prompt.BrowseDirectories(".")
	if selectedDir == "" {
		fmt.Println("Installation aborted.")
		return
	}
	err := os.MkdirAll(selectedDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", selectedDir, err)
		return
	}
	userDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error opening userhomedir: %v\n", err)
		return
	}
	smoothcloudDir := filepath.Join(userDir, ".smoothcloud")
	err = os.MkdirAll(smoothcloudDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", selectedDir, err)
		return
	}
	fmt.Println(smoothcloudDir)
	err = json.SaveJSON(smoothcloudDir + "/config.json", json.MainConfig{WorkingDirectory: selectedDir})
	if err != nil {
		fmt.Printf("Error saving global config: %v\n", err)
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
	databaseType := prompt.InputWithSelect("Which database type do you want to use for the cloud?", []string{"SQLITE", "MARIADB", "MONGODB"})
	var databaseConfig interface{}
	switch databaseType {
	case "MARIADB":
		databaseConfig = json.MariaDBDatabaseConfig{
			Host: prompt.Input("Enter the host of your database server", "127.0.0.1"),
			Port: prompt.InputPort("Enter the port of your database server", "3306"),
			Database: prompt.Input("Enter the database which you want to use", "smoothcloud"),
			Username: prompt.Input("Enter the username of your database user", ""),
			Password: prompt.Input("Enter the password of your database user", ""),
			Prefix: "smoothcloud_",
		}
	case "MONGODB":
		databaseConfig = json.MongoDBDatabaseConfig{
			Host: prompt.Input("Enter the host of your database server", "127.0.0.1"),
			Port: prompt.InputPort("Enter the port of your database server", "3306"),
			Database: prompt.Input("Enter the database which you want to use", "smoothcloud"),
			Username: prompt.Input("Enter the username of your database user", ""),
			Password: prompt.Input("Enter the password of your database user", ""),
			Prefix: "smoothcloud_",
		}
	default:
		databaseConfig = json.SQLiteDatabaseConfig{
			Filename: "sqlite.db",
			Prefix: "smoothcloud_",
		}
	}
	language := prompt.InputWithSelect("Enter the language for the cloud", []string{"en_US", "de_DE"})
	host := prompt.Input("Enter a IPv4 address which the cloud should use", "127.0.0.1")
	port := prompt.InputPort("Which port should the cloud use?", "6042")
	memory := prompt.InputInteger("What is the maximum amount of storage the cloud should use (in MB)?", "4096")
	config := json.Config{
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
