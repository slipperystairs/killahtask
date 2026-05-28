package cmd

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/slipperystairs/killahtask/cowsay"
	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var cow bool

type User struct {
	Username *user.User
	Filename string
	Filepath string
}

// TODO => I think we can make this better by using a hash map of sorts.
// TODO => map[string]bool instead of records [][]string
func uniqueDescription(task string, records [][]string) bool {
	for _, record := range records[1:] {
		fmt.Println(record)
		if task == record[1] {
			return false
		}
	}

	return true
}

func checkCowsay(message string, split bool) {
	lines := []string{message}

	if !cow {
		fmt.Printf("%s\n", message)
	} else {
		if split {
			lines = strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")
		}
		cowsay.CowSay(lines)
	}
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
