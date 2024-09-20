package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TokenCountRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type TokenCountResponse struct {
	Tokens int `json:"tokens"`
}

func CountTokens(url string, model Model, prompt string) (int, error) {
	tokenURL := url + "/api/tokens"

	requestBody, err := json.Marshal(TokenCountRequest{
		Model:  model.Name,
		Prompt: prompt,
	})
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(tokenURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var tokenResponse TokenCountResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return 0, err
	}

	return tokenResponse.Tokens, nil
}
