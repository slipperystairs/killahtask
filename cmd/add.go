package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

func Now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

var addCommand = &cobra.Command{
	Use:     "add",
	Short:   "Adds a new item",
	Aliases: []string{"a"},
	Long:    `This command will add a item to your list`,
	Run: func(cmd *cobra.Command, args []string) {
		var earlyExit bool = false
		if len(args) == 0 {
			PrintMsg("add", "add_none")
			earlyExit = true
		} else if len(args) > 1 { // A shitty way of making the user wrap their command in double quotes lol
			PrintMsg("add", "add_to_many")
			earlyExit = true
		}

		if earlyExit {
			os.Exit(1)
		}

		var description string = strings.TrimSpace(args[0])
		var fileExist bool = true
		// os.OpenFile doesn't have a way of letting us know if the file already exist.
		_, err := os.Stat(CurrentUser.Filepath)
		if os.IsNotExist(err) {
			fileExist = false
		} else {
			task.CheckError(err)
		}
		fmt.Printf("fileExist: %t\n", fileExist)

		file, err := task.LoadFile(CurrentUser.Filepath)
		// File will get closed even in the event of an error.
		defer task.CloseFile(file)
		task.CheckError(err)

		if !fileExist {
			records := [][]string{
				{"task_id", "description", "created", "completed"},
				{"0", description, Now(), "false"},
			}
			fmt.Println("in the made if")
			err := task.WriteCSV(file, records)
			task.CheckError(err)
		} else {
			csvReader := csv.NewReader(file)
			records, err := csvReader.ReadAll()
			task.CheckError(err)

			if len(records) > 0 {
				// Get the last task_id used and increment it by one.
				lastId, err := strconv.Atoi(records[len(records)-1][0])
				task.CheckError(err)
				newId := strconv.Itoa(lastId + 1)

				if !uniqueDescription(description, records) {
					fmt.Printf("Task description isn't unique! \"%s\" already exist.\n", description)
					os.Exit(1)
				}

				// Append the new record to the end of the slice.
				records = append(records, []string{newId, description, Now(), "true"})
				err = task.WriteCSV(file, records) // Re-write the file with the new records
				task.CheckError(err)
			}
		}

		fmt.Printf("Task \"%s\" added successfully!\n", description)
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
