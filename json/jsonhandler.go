package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SaveJSON(filePath string, data interface{}) error {
	formattedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error while formatting file: %w", err)
	}
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error while creating file: %w", err)
	}
	defer file.Close()
	_, err = file.Write(formattedJSON)
	if err != nil {
		return fmt.Errorf("error while writing file: %w", err)
	}
	return nil
}

func LoadJSON(filePath string, result interface{}) error {
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error while opening file: %w", err)
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error while reading file: %w", err)
	}
	err = json.Unmarshal(fileContent, result)
	if err != nil {
		return fmt.Errorf("error while parsing file: %w", err)
	}
	return nil
}

func LoadJSONFromURL(url string, result interface{}) error {
	fmt.Println(url)
	
	// HTTP-GET-Anfrage ausführen
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while fetching URL: %w", err)
	}
	defer resp.Body.Close()

	// Überprüfen, ob der HTTP-Statuscode erfolgreich ist
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	// Inhalt der Antwort lesen
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while reading response body: %w", err)
	}

	// JSON-Daten unmarshallen
	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("error while parsing JSON: %w", err)
	}

	return nil
}

func AddKeyValue(filePath string, key string, value interface{}) error {
	var data map[string]interface{}
	err := LoadJSON(filePath, &data)
	if err != nil && os.IsNotExist(err) {
		data = make(map[string]interface{})
	} else if err != nil {
		return fmt.Errorf("error while reading file: %w", err)
	}
	data[key] = value
	return SaveJSON(filePath, data)
}

func UpdateKeyValue(filePath string, key string, value interface{}) error {
	var data map[string]interface{}
	err := LoadJSON(filePath, &data)
	if err != nil {
		return fmt.Errorf("error while reading file: %w", err)
	}
	if _, exists := data[key]; !exists {
		return fmt.Errorf("key '%s' does'nt exist", key)
	}
	data[key] = value
	return SaveJSON(filePath, data)
}

func RemoveKey(filePath string, key string) error {
	var data map[string]interface{}
	err := LoadJSON(filePath, &data)
	if err != nil {
		return fmt.Errorf("error while reading file: %w", err)
	}
	if _, exists := data[key]; !exists {
		return fmt.Errorf("key '%s' does'nt exist", key)
	}
	delete(data, key)
	return SaveJSON(filePath, data)
}
