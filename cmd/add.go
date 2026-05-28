package cmd

import (
	"encoding/csv"
	"errors"
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
	Long:    `Adds a task to your list of TODOs by creating or updating your ~/killahtask_<username>.csv file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Missing task description")
		}

		description := strings.TrimSpace(strings.Join(args, " "))
		if len(description) > 50 {
			return errors.New("Description length is limited to 50 characters")
		}

		successMsg := fmt.Sprintf("\"%s\" added successfully!", description)
		// os.OpenFile doesn't have a way of letting us know if the file already exist.
		fileInfo, err := os.Stat(CurrentUser.Filepath)
		fileExists := err == nil

		// We check the file size because I ran into a bug during development where my file existed but it was empty.
		// Treat it like a new file when the user falls into this case for whatever reason.
		// TODO => Do more testing around this area to see if we can clean it up a bit.
		// TODO => If not, lets add some more comments to explain our decision here because it reads a little weird.
		hasData := fileExists && fileInfo.Size() > 0
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		file, err := task.LoadFile(CurrentUser.Filepath)
		if err != nil {
			return err
		}
		// File will close even in the event of an error.
		defer task.CloseFile(file)

		if !fileExists || !hasData {
			records := [][]string{
				{"task_id", "description", "created", "completed"},
				{"0", description, Now(), "false"},
			}

			err := task.WriteCSV(file, records)
			if err != nil {
				return err
			}
		} else {
			newId := "0"
			csvReader := csv.NewReader(file)
			records, err := csvReader.ReadAll()
			// TODO => This is where we would check if our map has records before call uniqueDescription()
			if err != nil {
				return err
			}

			// It is possible for the user to delete the only record and be left with just the headers
			// If that is the case we reset the ID to 0.
			if len(records) > 1 {
				// Get the last task_id used and increment it by one.
				lastId, err := strconv.Atoi(records[len(records)-1][0])
				if err != nil {
					return err
				}
				newId = strconv.Itoa(lastId + 1)
			}

			if !uniqueDescription(description, records) {
				errStr := "Task description isn't unique!\n\"" + description + "\" already exists.\n"
				return errors.New(errStr)
			}

			// Append the new record to the end of the slice.
			records = append(records, []string{newId, description, Now(), "false"})
			csvErr := task.WriteCSV(file, records) // Re-write the file with the new records
			if csvErr != nil {
				return csvErr
			}
		}

		checkCowsay(successMsg, false)
		return nil
	},
}

func init() {
	addCommand.PersistentFlags().BoolVar(&cow, "cowsay", false, "Display output using Cowsay")
	rootCmd.AddCommand(addCommand)
}
