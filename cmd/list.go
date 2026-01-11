package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List the items in your list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is where we would list your shits")
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
