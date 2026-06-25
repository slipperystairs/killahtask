package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/slipperystairs/killahtask/cowsay"
	"github.com/slipperystairs/killahtask/chodesay"
	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var cow bool
var chode bool
var descriptions = make(map[string]bool)

type User struct {
	Username *user.User
	Filename string
	Filepath string
}

func uniqueDescription(task string) bool {
	if descriptions[task] {
		return false
	}
	return true
}

func checkCowsay(message string, split bool) {
	lines := []string{message}

	if cow {
		if split {
			lines = strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")
		}
		cowsay.CowSay(lines)
	} else if chode {
		if split {
			lines = strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")
		}
		chodesay.ChodeSay(lines)
	} else {
			fmt.Printf("%s\n", message)
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
	task.MaybeDieAboutIt(err)
	rootCmd.PersistentFlags().Bool("cowsay", false, "Display output using cowsay")
	rootCmd.PersistentFlags().Bool("chodesay", false, "Display output as a chode")

	CurrentUser = User{
		Username: currUser,
		Filename: "killahtask_" + currUser.Username + ".csv",
		Filepath: filepath.Join(currUser.HomeDir, "killahtask_"+currUser.Username+".csv"),
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}
