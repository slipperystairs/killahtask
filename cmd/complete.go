package cmd

import (
	"encoding/csv"
	"errors"

	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var completeCommand = &cobra.Command{
	Use:     "complete",
	Short:   "Completes an item on the list",
	Aliases: []string{"c"},
	Long:    `Completes and item in your list by marking the "Completed" column as true.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Task ID is missing!")
		} else if len(args) > 1 {
			return errors.New("Too many arguments!")
		}
		cmd.SilenceUsage = true

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

		found := false
		if len(records) > 0 {
			for _, rec := range records[1:] {
				if rec[0] == args[0] {
					rec[3] = "true"
					found = true
					break
				}
			}
		}

		if !found {
			return errors.New("Task ID could not be found.")
		}

		csvErr := task.WriteCSV(file, records)
		if csvErr != nil {
			return csvErr
		}

		msg := "ID " + args[0] + " was marked as complete!"
		checkCowsay(msg, false)

		return nil
	},
}

func init() {
	completeCommand.PersistentFlags().BoolVar(&cow, "cowsay", false, "Display output using Cowsay")
	completeCommand.PersistentFlags().BoolVar(&chode, "chodesay", false, "Display output as a chode")
	rootCmd.AddCommand(completeCommand)
}
