package cmd

import (
	"os"

	"github.com/silvergama/unico/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "unico",
	Short: "unico - Evaluation process",
	Long:  "unico - Evaluation process",
}

func Execute() {
	logger.SetupLogger()
	if err := RootCmd.Execute(); err != nil {
		logrus.Panicln(err)
		os.Exit(-1)
	}
}
