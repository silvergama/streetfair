package cmd

import (
	"github.com/silvergama/streetfair/api"
	"github.com/silvergama/streetfair/repository"
	"github.com/spf13/cobra"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api",
		Short: "Initialize the API",
		Long:  "Initialize the API",
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
