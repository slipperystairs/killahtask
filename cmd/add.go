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
		earlyExit := false
		if len(args) == 0 {
			command = "add"
			PrintMsg(&command, "add_none")
			earlyExit = true
		} else if len(args) > 1 { // A shitty way of making the user wrap their command in double quotes lol
			command = "add"
			PrintMsg(&command, "add_to_many")
			earlyExit = true
		}

		if earlyExit {
			os.Exit(1)
		}

		description := strings.TrimSpace(args[0])
		successMsg := fmt.Sprintf("Task \"%s\" added successfully!\n", description)
		// os.OpenFile doesn't have a way of letting us know if the file already exist.
		fileInfo, err := os.Stat(CurrentUser.Filepath)
		fileExists := err == nil
		// We check the file size because I ran into a bug during development where my file existed but it was empty.
		// Treat it like a new file we the user falls into this case for whatever reason.
		hasData := fileExists && fileInfo.Size() > 0
		if err != nil && !os.IsNotExist(err) {
			task.CheckError(err)
		}

		file, err := task.LoadFile(CurrentUser.Filepath)
		// File will get closed even in the event of an error.
		defer task.CloseFile(file)
		task.CheckError(err)

		if !fileExists || !hasData {
			records := [][]string{
				{"task_id", "description", "created", "completed"},
				{"0", description, Now(), "false"},
			}
			err := task.WriteCSV(file, records)
			task.CheckError(err)
			fmt.Printf("%s", successMsg)
		} else {
			newId := "0"
			csvReader := csv.NewReader(file)
			records, err := csvReader.ReadAll()
			task.CheckError(err)

			// It is possible for the user to delete the only record and be left with just the headers
			// If that is the case we reset the ID to 0.
			if len(records) > 1 {
				// Get the last task_id used and increment it by one.
				lastId, err := strconv.Atoi(records[len(records)-1][0])
				task.CheckError(err)
				newId = strconv.Itoa(lastId + 1)
			}

			if !uniqueDescription(description, records) {
				fmt.Printf("Task description isn't unique! \"%s\" already exist.\n", description)
				os.Exit(1)
			}

			// Append the new record to the end of the slice.
			records = append(records, []string{newId, description, Now(), "false"})
			err = task.WriteCSV(file, records) // Re-write the file with the new records
			task.CheckError(err)
			fmt.Printf("%s", successMsg)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
