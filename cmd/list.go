package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/slipperystairs/killahtask/task"
	"github.com/spf13/cobra"
)

var headerMap = map[string]string{
	"task_id":     "ID",
	"description": "Description",
	"created":     "Created",
	"completed":   "Completed",
}

func PrintRecords(w *tabwriter.Writer, records [][]string) {
	header := records[0]
	if len(header) != 4 {
		panic("Header must have 4 columns.")
	}

	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
		headerMap[header[0]],
		headerMap[header[1]],
		headerMap[header[2]],
		headerMap[header[3]],
	)

	for _, rec := range records[1:] {
		if len(rec) != 4 {
			panic("Records must have 4 columns.")
		}
		fmt.Fprintln(w, strings.Join(rec, "\t"))
	}
}

// todo =>  - If sub flag is passed the show all records
// todo =>  - Else filter out task that are not marked as "complete" into a new slice
// todo =>  - Don't forget to use timediff to print out the time in words
// todo =>  - Close the file at this point
// todo => Figure out how to use the text/tabwriter package to print the records into a table
var listCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List the items in your list",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := task.LoadFile(CurrentUser.Filepath)
		task.CheckError(err)
		defer task.CloseFile(file)

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		task.CheckError(err)

		if len(records) == 0 {
			fmt.Println("No task to show! Try adding a task by running killahtask add \"my task\"")
		} else {
			w := tabwriter.NewWriter(os.Stdout, 0, 4, 3, ' ', 0)
			defer w.Flush()
			PrintRecords(w, records)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
