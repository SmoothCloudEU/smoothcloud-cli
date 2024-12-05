package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while fetching URL: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while reading response body: %w", err)
	}
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

func AddNestedKeyValue(filePath string, keyPath string, value interface{}) error {
    var data map[string]interface{}
    err := LoadJSON(filePath, &data)
    if err != nil && os.IsNotExist(err) {
        data = make(map[string]interface{})
    } else if err != nil {
        return fmt.Errorf("error while reading file: %w", err)
    }

    // Schlüsselpfad in eine verschachtelte Struktur einfügen
    if err := setNestedKey(data, keyPath, value); err != nil {
        return fmt.Errorf("error while setting nested key: %w", err)
    }

    return SaveJSON(filePath, data)
}

func setNestedKey(data map[string]interface{}, keyPath string, value interface{}) error {
    keys := strings.Split(keyPath, ".") // Schlüssel in einzelne Teile zerlegen
    current := data

    for i, key := range keys {
        if i == len(keys)-1 {
            // Letzter Schlüssel -> Wert setzen
            current[key] = value
        } else {
            // Schlüssel existiert nicht -> neue Map erstellen
            if _, exists := current[key]; !exists {
                current[key] = make(map[string]interface{})
            }

            // Typprüfung: Ist der aktuelle Schlüssel eine Map?
            nested, ok := current[key].(map[string]interface{})
            if !ok {
                return fmt.Errorf("key '%s' is not a map", key)
            }
            current = nested
        }
    }

    return nil
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
