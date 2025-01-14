package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "auth",
	Short: "authentication server binary",
}

func init() {
	rootCmd.AddCommand(serveRestCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
