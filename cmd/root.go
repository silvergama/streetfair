package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "unico",
	Short: "unico - Evaluation process",
	Long:  "unico - Evaluation process",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Panic(err.Error())
		os.Exit(-1)
	}
}
