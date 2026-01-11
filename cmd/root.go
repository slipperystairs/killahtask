package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var flagValue string

var rootCmd = &cobra.Command{
	Use:   "killahtask",
	Short: "Killah Task is a todo CLI tool.",
	Long:  `A todo task tool that does things`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Len %d\n", len(args))
		// if item != "" {
		// 	fmt.Printf("Your item is %s \n", item);
		// }
		
		// if id != "" {
		// 	fmt.Printf("This is where I would")
		// }
		
		if len(args) == 0 {
			fmt.Println("No commands were passed to the killah...see below")
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("list", "l", false, "Prints all the items in your todo list")
	rootCmd.PersistentFlags().StringVarP(&flagValue, "add", "a", "", "Item to be added to your list")
	// TODO: We will have to figure this one out.
	rootCmd.PersistentFlags().StringVarP(&flagValue, "complete", "c", "", "Completes an item on your list")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
