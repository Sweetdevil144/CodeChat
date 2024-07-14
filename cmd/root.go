package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "CodeChat",
    Short: "CodeChat App is a terminal-based chat application",
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.AddCommand(chatCmd)
    rootCmd.AddCommand(getfs)
}
