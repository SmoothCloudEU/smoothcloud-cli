package prompt

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)


func BrowseDirectories(baseDir string) string {
	currentDir := baseDir
	for {
		dirs, err := GetSubdirectories(currentDir)
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
			return BrowseDirectories(result)
		}
		if result == ".." {
			currentDir = filepath.Dir(currentDir)
			continue
		}
		currentDir = filepath.Join(currentDir, result)
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