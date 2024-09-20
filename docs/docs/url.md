# URL

The `webassess url` performs an assessment of the underlying HTML code found at a URL.

## Usage

```bash
webassess url --target http://example.com --output json

```

### Help Text

```bash
$ webassess url -h
Perform a URL content assessment against a URL target

Usage:
  webassess url [flags]

Flags:
  -h, --help            help for url
      --target string   URL target to perform web AI assessment against

Global Flags:
  -d, --allow-download        Allow downloading of models from internet if not already available
  -m, --ollama-model string   Ollama model and version to use for assessment (default "qwen2.5:0.5b")
  -u, --ollama-url string     URL for Ollama service
  -o, --output string         Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string    Path to output file. If blank, will output to STDOUT
  -q, --quiet                 Suppress output
  -v, --verbose               Verbose output
```
