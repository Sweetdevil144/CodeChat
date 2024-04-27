package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var start string

// startCmd is the root command here which will begin the CLI
var startCmd = &cobra.Command{
	Use:   "CodeChat is a CLI Chat Application for Developers",
	Short: "This is short description for start command of CodeChat",
	Long:  "This is Long description for start command of CodeChat",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start flag is : ", start)
	},
	TraverseChildren: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := startCmd.Execute()
	if err != nil {
		os.Exit(2)
	}
}

func init() {
	startCmd.Flags().StringVarP(&start, "start", "s", "default_start_command_value", "Start the Chat Application")
	err := viper.BindPFlag("start", startCmd.Flags().Lookup("start"))
	if err != nil {
		fmt.Println("Error in retrieving flag in startCmd")
		return
	}
}
