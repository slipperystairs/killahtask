package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var (
	showAll bool
	buf     strings.Builder
)

var headerMap = map[string]string{
	"task_id":     "ID",
	"description": "Description",
	"created":     "Created",
	"completed":   "Completed",
}

// Parse the time string using the RFC3339 standard i.e., "2026-01-020T15:04:05Z07:00".
// Convert into local time and returns the difference in human readable format.
func timeDiff(rec string) string {
	parsed, err := time.Parse(time.RFC3339, rec)
	if err != nil {
		return "invalid time"
	}
	localTime := parsed.Local()

	return timediff.TimeDiff(localTime)
}

// The conditions where panic is called should never happen, but we might as well be prepared.
func buildTable(w *tabwriter.Writer, records [][]string, all bool) error {
	header := records[0]
	if len(header) != 4 {
		return errors.New("Invalid task file format")
	}

	if !all {
		fmt.Fprintf(w, "%s\t%s\t%s\t\n",
			headerMap[header[0]],
			headerMap[header[1]],
			headerMap[header[2]],
		)
	} else {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			headerMap[header[0]],
			headerMap[header[1]],
			headerMap[header[2]],
			headerMap[header[3]],
		)
	}

	for _, rec := range records[1:] {
		if len(rec) != 4 {
			return errors.New("Invalid task record format")
		}

		diff := timeDiff(rec[2])
		if diff == "invalid time" {
			return errors.New("Invalid time format passed to the timediff parser")
		}

		// If the sub flag is passed then we show the "Completed" column and all of the task
		if !all && rec[3] != "true" {
			fmt.Fprintf(w, "%s\t%s\t%s\n",
				rec[0],
				rec[1],
				diff,
			)
		} else if all {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
				rec[0],
				rec[1],
				diff,
				rec[3],
			)
		}
	}

	return nil
}

var listCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List the items in your list",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := task.LoadFile(CurrentUser.Filepath)
		if err != nil {
			return err
		}
		defer task.CloseFile(file)

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return err
		}

		if len(records) == 0 {
			str := "No task to show! Try adding a task by running killahtask add \"my task\""
			return errors.New(str)
		}

		w := tabwriter.NewWriter(&buf, 0, 4, 5, ' ', 0)
		tabwriterErr := buildTable(w, records, showAll)
		w.Flush()

		if tabwriterErr != nil {
			return tabwriterErr
		}

		split := false
		if cow {
			split = true
		}
		checkCowsay(buf.String(), split)

		return nil
	},
}

func init() {
	listCommand.Flags().BoolVarP(&showAll, "all", "a", false, "Shows all flag task items (alias: -a)")
	listCommand.PersistentFlags().BoolVar(&cow, "cowsay", false, "Display output using Cowsay")
	rootCmd.AddCommand(listCommand)
}
