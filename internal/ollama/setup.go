package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

type Model struct {
	Name              string `json:"name"`
	ModifiedAt        string `json:"modified_at"`
	Size              int64  `json:"size"`
	ContextWindowSize int    `json:"context_window_size"`
}

func GetAvailableOllamaModels(url string) ([]Model, error) {
	tagsURL := url + "/api/tags"

	resp, err := http.Get(tagsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			if err == nil {
				err = fmt.Errorf("error closing response body: %w", closeErr)
			} else {
				fmt.Printf("error closing response body: %v\n", closeErr)
			}
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result struct {
		Models []Model `json:"models"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return result.Models, nil
}

func ModelReady(url string, modelName string) bool {
	models, err := GetAvailableOllamaModels(url)
	if err != nil {
		return false
	}

	for _, model := range models {
		if model.Name == modelName {
			return true
		}
	}

	return false
}

func GetModel(url string, modelName string) (Model, error) {
	models, err := GetAvailableOllamaModels(url)
	if err != nil {
		return Model{}, err
	}

	for _, model := range models {
		if model.Name == modelName {
			return model, nil
		}
	}

	return Model{}, fmt.Errorf("model not found")
}

func DownloadOllamaModel(modelName string, url string) error {
	pullURL := url + "/api/pull"

	requestBody, err := json.Marshal(map[string]string{
		"name": modelName,
	})
	if err != nil {
		return fmt.Errorf("failed to create request body: %v", err)
	}

	resp, err := http.Post(pullURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			if err == nil {
				err = fmt.Errorf("error closing response body: %w", closeErr)
			} else {
				fmt.Printf("error closing response body: %v\n", closeErr)
			}
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create a channel to send progress updates
	progress := make(chan string)

	// Start a goroutine to read and process the response body
	go func() {
		defer close(progress)

		decoder := json.NewDecoder(resp.Body)
		for {
			var result map[string]interface{}
			if err := decoder.Decode(&result); err != nil {
				if err == io.EOF {
					break
				}
				progress <- fmt.Sprintf("Error decoding response: %v", err)
				return
			}

			if status, ok := result["status"].(string); ok {
				progress <- status
			}
		}
	}()

	// Print progress updates
	for status := range progress {
		fmt.Println(status)
	}

	return nil
}

func IsOllamaRunning(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			if err == nil {
				err = fmt.Errorf("error closing response body: %w", closeErr)
			} else {
				fmt.Printf("error closing response body: %v\n", closeErr)
			}
		}
	}()
	return resp.StatusCode == http.StatusOK
}

func StartOllama() error {
	cmd := exec.Command("ollama", "serve")
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start Ollama: %v", err)
	}

	// Wait for Ollama to start
	time.Sleep(5 * time.Second)
	return nil
}
