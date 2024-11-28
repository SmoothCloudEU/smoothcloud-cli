package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	host := strings.TrimSpace(Input("Which IPv4-adresse should use the cloud?"))
	port := strings.TrimSpace(Input("Which port should use the cloud?"))
	memory := strings.TrimSpace(Input("What is the maximum amount of storage the cloud should use?"))
	fmt.Printf("Cloud has been successfully installed into %s with Host %s:%s!\nMemory: %s", directory, host, port, memory)
}

func Input(question string) string {
	fmt.Printf("\n%s\n", question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while read input:", err)
		return ""
	}
	return input
}
