package command

import (
	"fmt"
	"os"
)

func Install(directory string) {
	fmt.Printf("Installing cloud...\n")
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		fmt.Printf("Error with creating directory %s.\n", directory)
		return
	}
	directories := []string{"/groups", "/groups/proxies", "/groups/lobbies", "/groups/servers", "/storage", "/templates"}
	for _, value := range directories {
		dir := directory + value
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error with creating directory %s.\n", dir)
			os.Exit(1)
			return
		}
	}
	fmt.Printf("Cloud has been successfully installed into %s!\n", directory)
}
