package cmd

import (
	"github.com/silvergama/streetfair/importdb"
	"github.com/silvergama/streetfair/repository"
	"github.com/spf13/cobra"
)

var (
	importCommand = &cobra.Command{
		Use:   "import",
		Short: "Import data to database",
		Long:  "Import data to database from file",
		Args:  cobra.ExactArgs(1),
		RunE:  importExecute,
	}
)

func init() {
	RootCmd.AddCommand(importCommand)
}

func importExecute(cmd *cobra.Command, args []string) error {
	err := repository.Setup()
	if err != nil {
		return err
	}
	imp := importdb.NewImportCSV()
	err = imp.ImportFromCSV(args[0])
	if err != nil {
		return err
	}
	return nil
}
