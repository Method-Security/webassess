package cmd

import (
	"github.com/Method-Security/webassess/internal/url"
	"github.com/spf13/cobra"
)

// InitURLAssess initializes the url command for the webassess CLI. This command is used to perform dynamic analysis
// of the contents and files hosted on a URL.
func (a *WebAssess) InitURLAssess() {
	urlCmd := &cobra.Command{
		Use:   "url",
		Short: "Perform a URL content assessment against a URL target",
		Long:  `Perform a URL content assessment against a URL target`,
		Run: func(cmd *cobra.Command, args []string) {
			target, err := cmd.Flags().GetString("target")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			report := url.PerformURLAssess(cmd.Context(), target, a.RootFlags.OllamaModel)

			a.OutputSignal.Content = report
		},
	}

	urlCmd.Flags().String("target", "", "URL target to perform web AI assessment against")

	a.RootCmd.AddCommand(urlCmd)
}
