package cmd

import (
	"encoding/csv"
	"fmt"

	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var completeCommand = &cobra.Command{
	Use:     "complete",
	Short:   "Completes an item on the list",
	Aliases: []string{"c"},
	Long:    `This command will complete and item on your list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			command = "complete"
			PrintMsg(&command, "comp_too_many")
		} else {
			file, err := task.LoadFile(CurrentUser.Filepath)
			defer task.CloseFile(file)
			task.CheckError(err)

			csvReader := csv.NewReader(file)
			records, err := csvReader.ReadAll()
			task.CheckError(err)

			found := false
			for _, rec := range records[1:] {
				if rec[0] == args[0] {
					rec[3] = "true"
					found = true
					break
				}
			}

			if !found {
				PrintMsg(nil, "unknown_id")
			} else {
				task.WriteCSV(file, records)
				fmt.Printf("Task %s was marked as complete!\n", args[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCommand)
}
