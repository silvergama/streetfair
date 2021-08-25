package cmd

import (
	"github.com/silvergama/unico/api"
	"github.com/silvergama/unico/repository"
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
	if err := repository.Setup(); err != nil {
		return err
	}

	return api.Setup()
}
