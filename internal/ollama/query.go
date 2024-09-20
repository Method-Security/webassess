package ollama

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ollama/ollama/api"
)

var ErrContextLengthExceeded = errors.New("context length exceeded")

// IsContextLengthError checks if an error is due to context length exceeding the model's limit.
func IsContextLengthError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, ErrContextLengthExceeded) {
		return true
	}
	if strings.Contains(err.Error(), "context window exceeded") || strings.Contains(err.Error(), "too many tokens") {
		return true
	}
	return false
}

// QueryModel queries the specified model with the given prompt using the Ollama SDK.
func QueryModel(ctx context.Context, client *api.Client, model Model, prompt string) (string, error) {
	var result strings.Builder

	req := &api.GenerateRequest{
		Model:  model.Name,
		Prompt: prompt,
	}

	err := client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		result.WriteString(resp.Response)
		return nil
	})
	if err != nil {
		if IsContextLengthError(err) {
			return "", ErrContextLengthExceeded
		}
		return "", fmt.Errorf("failed to generate response: %v", err)
	}

	return result.String(), nil
}

type ModelPromptContentGenerator func(string) string

// ProcessContentRecursively processes the content recursively, splitting it if necessary.
// The input always gets the same prompt generator call to ensure the instructions are consistent across splits.
func ProcessContentRecursively(ctx context.Context, client *api.Client, model Model, input string, generator ModelPromptContentGenerator, combiner ModelPromptContentGenerator) (string, error) {
	content := generator(input)
	// Attempt to query the model
	response, err := QueryModel(ctx, client, model, content)
	if err != nil {
		if IsContextLengthError(err) && len(input) > 100 {
			// If context length is exceeded, split the content and process each half
			mid := len(input) / 2
			leftInput := input[:mid]
			rightInput := input[mid:]

			// Recursively process the left half
			leftResult, errLeft := ProcessContentRecursively(ctx, client, model, leftInput, generator, combiner)
			if errLeft != nil {
				return "", errLeft
			}

			// Recursively process the right half
			rightResult, errRight := ProcessContentRecursively(ctx, client, model, rightInput, generator, combiner)
			if errRight != nil {
				return "", errRight
			}

			// Combine the results
			combinedOutput := leftResult + "\n\n" + rightResult
			combinedPrompt := combiner(combinedOutput)

			finalResult, err := QueryModel(ctx, client, model, combinedPrompt)
			if err != nil {
				return "", err
			}

			return finalResult, nil
		}
		// Return other errors
		return "", err
	}

	return response, nil
}
