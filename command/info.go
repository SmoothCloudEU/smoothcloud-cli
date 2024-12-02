package command

import (
	"fmt"

	"smoothcloudcli/json"
)

func Info() {
	var result map[string]interface{}
	json.LoadJSON("cloud/config.json", result)
	fmt.Println("-----------------Information-----------------")
	fmt.Println("SmoothCloud v1.0-indev")
	fmt.Println("xxxxxx@dev")
	fmt.Printf("Memory: %s\n", result)
	fmt.Println("-----------------Information-----------------")
}
