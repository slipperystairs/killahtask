package cmd

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var cow bool
type User struct {
	Username *user.User
	Filename string
	Filepath string
}

func uniqueDescription(task string, records [][]string) bool {
	for _, record := range records[1:] {
		if task == record[1] {
			return false
		}
	}

	return true
}

var CurrentUser User
var rootCmd = &cobra.Command{
	Use:   "killahtask",
	Short: "Killah Task is a todo CLI tool.",
	Long:  `A todo task tool that performs simple CRUD operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No commands were passed to the killah...see below")
			cmd.Help()
		}
	},
}

func init() {
	currUser, err := user.Current()
	task.CheckError(err)
	rootCmd.PersistentFlags().Bool("cowsay", false, "Display output using cowsay")

	CurrentUser = User{
		Username: currUser,
		Filename: "killahtask_" + currUser.Username + ".csv",
		Filepath: filepath.Join(currUser.HomeDir, "killahtask_"+currUser.Username+".csv"),
	}
}

func Execute() {
	err := rootCmd.Execute()
	task.CheckError(err)
}
