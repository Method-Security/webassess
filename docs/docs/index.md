# Capabilities

webassess provides a number of capabilities to performing AI based assessments of security resources.

- [URL](./url.md)

## Top Level Flags

webassess has several top level flags that can be used on any subcommand. These include:

```bash
Flags:
  -h, --help                  help for webassess
  -d, --allow-download        Allow downloading of models from internet if not already available
  -m, --ollama-model string   Ollama model and version to use for assessment (default "qwen2.5:0.5b")
  -u, --ollama-url string     URL for Ollama service
  -o, --output string         Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string    Path to output file. If blank, will output to STDOUT
  -q, --quiet                 Suppress output
  -v, --verbose               Verbose output
```

## Version Command

Run `webassess version` to get the exact version information for your binary

## Output Formats

For more information on the various output formats that are supported by webassess, see the [Output Formats](https://method-security.github.io/docs/output.html) page in our organization wide documentation.
