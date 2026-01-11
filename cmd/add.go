package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var addCommand = &cobra.Command{
	Use: "add",
	Short: "Adds a new item",
	Long: `This command will add a item to your list`,
	Run: func(cmd *cobra.Command, args[]string) {
		fmt.Println("This is where I will add your items")
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
