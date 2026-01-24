package cmd

import (
	"encoding/csv"
	"fmt"

	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete an item in your list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			command = "delete"
			PrintMsg(&command, "delete_none")
		} else if len(args) > 1 {
			PrintMsg(nil, "delete_too_many")
		} else {
			file, err := task.LoadFile(CurrentUser.Filepath)
			defer task.CloseFile(file)
			task.CheckError(err)

			csvReader := csv.NewReader(file)
			records, err := csvReader.ReadAll()
			task.CheckError(err)

			// Create the new records from scratch
			newRecords := [][]string{}
			for _, rec := range records {
				if rec[0] != args[0] {
					newRecords = append(newRecords, rec)
				}
			}

			if len(newRecords) == len(records) {
				PrintMsg(nil, "unknown_id")
			} else {
				err = task.WriteCSV(file, newRecords)
				task.CheckError(err)
				fmt.Println("Task removed successfully!")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCommand)
}
