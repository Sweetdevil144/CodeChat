package cmd

import (
    "CodeChat/internal/server"
    "github.com/spf13/cobra"
    "log"
)

var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Starts the chat server",
    Run: func(cmd *cobra.Command, args []string) {
        err := server.StartServer()
        if err != nil {
            log.Fatalf("Failed to start the server: %v", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(serverCmd)
}
