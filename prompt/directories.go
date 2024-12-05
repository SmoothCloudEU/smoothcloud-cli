package prompt

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)


func BrowseDirectories(baseDir string) string {
	// Hole den absoluten Pfad des Basisverzeichnisses
	currentDir, err := filepath.Abs(baseDir)
	if err != nil {
		fmt.Printf("Error determining absolute path: %v\n", err)
		return ""
	}

	for {
		// Unterverzeichnisse holen
		dirs, err := GetSubdirectories(currentDir)
		if err != nil {
			fmt.Printf("Error reading directories: %v\n", err)
			return ""
		}

		// Navigationsoptionen hinzufügen
		if currentDir != baseDir {
			dirs = append([]string{".."}, dirs...)
		}
		dirs = append(dirs, "Select this directory")
		dirs = append(dirs, "Create a directory")

		// Auswahl anzeigen
		prompt := promptui.Select{
			Label: fmt.Sprintf("Current directory: %s\nNavigate or select:", currentDir),
			Items: dirs,
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Error selecting directory: %v\n", err)
			return ""
		}

		// Aktionen basierend auf Auswahl
		switch result {
		case "Select this directory":
			absolutePath, err := filepath.Abs(currentDir)
			if err != nil {
				fmt.Printf("Error getting absolute path: %v\n", err)
				return ""
			}
			return absolutePath

		case "Create a directory":
			prompt := promptui.Prompt{
				Label:     "Name of the new directory",
				Default:   "cloud/",
				AllowEdit: true,
				Validate:  validateNonEmpty,
			}
			dirName, err := prompt.Run()
			if err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				return ""
			}
			newDir := filepath.Join(currentDir, dirName)
			if err := os.MkdirAll(newDir, os.ModePerm); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", newDir, err)
				return ""
			}
			return BrowseDirectories(newDir)

		case "..":
			currentDir = filepath.Dir(currentDir)

		default:
			currentDir = filepath.Join(currentDir, result)
		}
	}
}

func GetSubdirectories(path string) ([]string, error) {
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