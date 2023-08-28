package main

import (
	"os"

	"github.com/mjevans93308/platformscience/cmd"
	"github.com/mjevans93308/platformscience/util/logs"
)

func main() {
	rootCmd := cmd.Execute()
	slog := logs.NewSlog()
	if err := rootCmd.Execute(); err != nil {
		slog.Errorf("Error while processing files: %s", err)
		os.Exit(1)
	}
}
