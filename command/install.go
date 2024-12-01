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
	fmt.Println("Installing cloud...\n")

	// Verzeichnisauswahl mit Navigationsfunktion
	selectedDir := browseDirectories(".")
	if selectedDir == "" {
		fmt.Println("Installation aborted.")
		return
	}

	// Basisverzeichnis und Unterverzeichnisse erstellen
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

	// Konfiguration einlesen
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

	configPath := filepath.Join(selectedDir, "config.json")
	err = json.SaveJSON(configPath, config)
	if err != nil {
		fmt.Printf("Error saving configuration to %s: %v\n", configPath, err)
		return
	}

	fmt.Printf("\nCloud has been successfully installed into \"%s\"!\n", selectedDir)
}

// Verzeichnis-Navigation mit Optionen zum Wechseln und Zurückgehen
func browseDirectories(baseDir string) string {
	currentDir := baseDir

	for {
		// Liste der Verzeichnisse im aktuellen Verzeichnis abrufen
		dirs, err := getSubdirectories(currentDir)
		if err != nil {
			fmt.Printf("Error reading directories: %v\n", err)
			return ""
		}

		// Option zum Zurückgehen hinzufügen, wenn nicht im Basisverzeichnis
		if currentDir != baseDir {
			dirs = append([]string{".."}, dirs...)
		}

		// Option zum Verlassen hinzufügen
		dirs = append(dirs, "Select this directory")

		// Benutzeraufforderung
		prompt := promptui.Select{
			Label: fmt.Sprintf("Current directory: %s\nNavigate or select:", currentDir),
			Items: dirs,
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Error selecting directory: %v\n", err)
			return ""
		}

		// Verzeichnis auswählen
		if result == "Select this directory" {
			return currentDir
		}

		// Zurück ins übergeordnete Verzeichnis
		if result == ".." {
			currentDir = filepath.Dir(currentDir)
			continue
		}

		// In das gewählte Unterverzeichnis wechseln
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
