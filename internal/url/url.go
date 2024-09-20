package url

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	webassess "github.com/Method-Security/webassess/generated/go"
	"github.com/Method-Security/webassess/internal/ollama"
	"github.com/ollama/ollama/api"
)

func PerformURLAssess(ctx context.Context, target string, model ollama.Model) webassess.UrlReport {
	report := webassess.UrlReport{
		Target: target,
		Errors: []string{},
	}

	// Step 1: Fetch HTML content from the target URL
	htmlContent, err := fetchHTMLContent(target)
	if err != nil {
		report.Errors = append(report.Errors, fmt.Sprintf("Failed to fetch URL: %v", err))
		return report
	}

	// Step 2: Initialize Ollama client
	client, err := api.ClientFromEnvironment()
	if err != nil {
		report.Errors = append(report.Errors, fmt.Sprintf("Failed to create Ollama client: %v", err))
		return report
	}

	// Step 3: Process the content recursively
	finalOutput, err := ollama.ProcessContentRecursively(ctx, client, model, htmlContent, CreateHTMLAnalysisPrompt, CreateHTMLSynthesisPrompt)
	if err != nil {
		report.Errors = append(report.Errors, err.Error())
		return report
	}

	// Step 5: Set the final report
	report.Output = finalOutput

	return report
}

func fetchHTMLContent(target string) (string, error) {
	resp, err := http.Get(target)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %v", err)
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
		return "", fmt.Errorf("failed to fetch URL: status code %d", resp.StatusCode)
	}

	htmlContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(htmlContent), nil
}
