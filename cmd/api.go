package cmd

import (
	"github.com/silvergama/unico/api"
	"github.com/spf13/cobra"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api",
		Short: "api",
		Long:  "api",
		RunE:  apiExecute,
	}
)

func init() {
	RootCmd.AddCommand(apiCommand)
}

func apiExecute(cmd *cobra.Command, args []string) error {
	api.Setup()

	return nil
}
