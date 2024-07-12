package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "chat-app",
    Short: "Chat App is a terminal-based chat application using WebSockets",
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.AddCommand(chatCmd)
}
