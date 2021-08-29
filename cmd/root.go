package cmd

import (
	"os"

	"github.com/silvergama/streetfair/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "streetfair",
	Short: "streetfair - Centralizador de feiras de rua.",
	Long:  "streetfair - Centralizador de informações sobre feiras de rua na cidade de São Paulo.",
}

func Execute() {
	logger.SetupLogger()
	if err := RootCmd.Execute(); err != nil {
		logrus.Panicln(err)
		os.Exit(-1)
	}
}
