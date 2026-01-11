package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var id string

var completeCommand = &cobra.Command{
	Use: "complete",
	Short: "Completes an item on the list",
	Long: `This command will complete and item on your list`,
	Run: func(cmd *cobra.Command, args[] string) {
		fmt.Println("This is where I would complete your shits")
	},
}

func init() {
	rootCmd.AddCommand(completeCommand)
}
