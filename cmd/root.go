package cmd

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var Cow bool
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

func PrintMsg(command *string, msgCase string) {
	switch msgCase {
	case "missing_task":
		fmt.Println("Missing task description.")
	case "comp_none", "delete_none":
		fmt.Println("Task ID is missing!")
	case "comp_too_many":
		fmt.Println("Too many arguments passed to the \"complete\" command.")
	case "unknown_id":
		fmt.Println("Task ID could not be found.")
	case "delete_too_many":
		fmt.Println("Too many arguments passed to the \"delete\" command.")
	}

	if command != nil {
		switch *command {
		case "add":
			fmt.Println("Usage: killahtask add \"my description\"")
		case "delete", "complete":
			fmt.Printf("Usage: killahtask %s <task_id>\n", *command)
		case "list":
			fmt.Println("Usage: killahtask list (optional: --all, -a)")
		}
	}
}

var command string
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

	rootCmd.PersistentFlags().Bool("cowsay", false, "Use Cowsay to display output")

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
