package command

import (
	"bufio"
	"fmt"
	"os"
	"smoothcloudcli/json"
	"strconv"
	"strings"
)

type Config struct {
	Language string `json:"language"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Memory   int    `json:"memory"`
}

func Install() {
	fmt.Printf("Installing cloud...\n\n")
	directory := Input("In which directory should the cloud installed?")
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
	language := Input("Which language do you want to use for the cloud?")
	host := Input("Which IPv4-adresse should use the cloud?")
	port, err := strconv.Atoi(Input("Which port should use the cloud?"))
	if err != nil {
		fmt.Printf("Error with parsing port into integer")
		return
	}
	memory, err := strconv.Atoi(Input("What is the maximum amount of storage the cloud should use?"))
	if err != nil {
		fmt.Printf("Error with parsing port into integer")
		return
	}
	fmt.Printf("\nCloud has been successfully installed into \"%s\"!", directory)
	json.SaveJSON(directory+"config.json", Config{Language: language, Host: host, Port: port, Memory: memory})
}

func Input(question string) string {
	fmt.Printf("\n%s\n", question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error while read input: %s\n", err)
		return ""
	}
	return strings.TrimSpace(input)
}
