# Adding a new capability

By design, webassess has a unique command for every assessment. To add another capability:

1. Add a file to `cmd/` that corresponds to the sub-command name you'd like to add to the `webassess` CLI
2. You can use `cmd/url.go` as a template
3. Your file needs to be a member function of the `webassess` struct and should be of the form `Init<cmd>Command`
4. Add a new member to the `webassess` struct in `cmd/root.go` that corresponsds to your command name. Remember, the first letter must be capitalized.
5. Call your `Init` function from `main.go`
6. Add logic to your commands runtime and put it in its own package within `internal` (e.g., `internal/url`)
