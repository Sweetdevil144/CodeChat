package cmd

import (
    "CodeChat/internal/client"
    "github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
    Use:   "chat",
    Short: "Starts the chat client",
    RunE: func(cmd *cobra.Command, args []string) error {
        return client.StartClient()
    },
}
