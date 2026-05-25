package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"

	"github.com/slipperystairs/killahtask/cowsay"
	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short: "Deletes and item.",
	Long:   "Delete an item in your list and recreates the items in the CSV file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Task ID is missing!")
		} else if len(args) > 1 {
			return errors.New("Too many arguments!")
		}

		file, err := task.LoadFile(CurrentUser.Filepath)
		if err != nil {
			return err
		}
		defer task.CloseFile(file)

		csvReader := csv.NewReader(file)
		records, err := csvReader.ReadAll()
		if err != nil {
			return err
		}

		// Create the new records from scratch
		newRecords := [][]string{}
		for _, rec := range records {
			if rec[0] != args[0] {
				newRecords = append(newRecords, rec)
			}
		}

		if len(newRecords) == len(records) {
			return errors.New("Task ID could not be found.")
		}

		csvError := task.WriteCSV(file, newRecords)
		if csvError != nil {
			return csvError
		}

		msg := "Task " + args[0] + " removed successfully!"
		if !cow {
			fmt.Printf("%s\n", msg)
		} else {
			lines := []string{msg}
			cowsay.CowSay(lines)
		}

		return nil
	},
}

func init() {
	deleteCommand.PersistentFlags().BoolVar(&cow, "cowsay", false, "Display output using cowsay")
	rootCmd.AddCommand(deleteCommand)
}
