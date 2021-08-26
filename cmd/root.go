package cmd

import (
	"os"

	"github.com/silvergama/streetfair/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "streetfair",
	Short: "streetfair - Evaluation process",
	Long:  "streetfair - Evaluation process",
}

func Execute() {
	logger.SetupLogger()
	if err := RootCmd.Execute(); err != nil {
		logrus.Panicln(err)
		os.Exit(-1)
	}
}
