package prompt

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

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

func InputWithEmpty(question string, defaultValue string) string {
	prompt := promptui.Prompt{
		Label:     question,
		Default:   defaultValue,
		AllowEdit: false,
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

func InputPort(input string, suggestion string) int {
	for {
		portStr := Input(input, suggestion)
		port, err := strconv.Atoi(portStr)
		if err != nil || port <= 0 || port > 65535 {
			fmt.Println("Invalid port number. Please enter a valid port (1-65535).")
			continue
		}
		return port
	}
}

func InputInteger(input string, suggestion string) int {
	for {
		memoryStr := Input(input, suggestion)
		memory, err := strconv.Atoi(memoryStr)
		if err != nil {
			fmt.Println("Invalid memory number.")
			continue
		}
		return memory
	}
}

func validateNonEmpty(input string) error {
	if input == "" {
		return fmt.Errorf("input cannot be empty")
	}
	return nil
}
