// Package config contains common configuration values that are used by the various commands and subcommands in the CLI.
package config

import (
	ollama "github.com/Method-Security/webassess/internal/ollama"
)

type RootFlags struct {
	Quiet       bool
	Verbose     bool
	OllamaURL   string
	OllamaModel ollama.Model
}
