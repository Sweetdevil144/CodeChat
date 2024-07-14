package cmd 

import (
    "github.com/spf13/cobra"
	"CodeChat/utils"
)

// func getFS() {
// 	utils.ShowTreeView()
// }

var getfs = &cobra.Command{
	Use: "fs",
	Short: "Display file structure of Package",
	RunE: func(cmd *cobra.Command, args []string) error {
		utils.ShowTreeView()
		return nil
	},

}