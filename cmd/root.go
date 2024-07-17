package cmd

import (
	"github.com/leslieleung/hotline/cmd/exec"
	"github.com/leslieleung/hotline/internal/ui"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "Hotline is a command line tool for executing LLM workflows.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(exec.Cmd)
	rootCmd.PersistentFlags().BoolVarP(&ui.Debug, "debug", "d", false, "enable debug output")
}
