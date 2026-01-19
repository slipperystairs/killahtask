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

var all bool

var headerMap = map[string]string{
	"task_id":     "ID",
	"description": "Description",
	"created":     "Created",
	"completed":   "Completed",
}

// The conditions where panic is called should never happen, but we might as well be prepared.
func printRecords(w *tabwriter.Writer, records [][]string, all bool) {
	fmt.Printf("all here: %t\n", all)
	header := records[0]
	if len(header) != 4 {
		panic("Header must have 4 columns.")
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
			panic("Records must have 4 columns.")
		}

		if all {
			fmt.Fprintln(w, strings.Join(rec, "\t"))
		} else if rec[3] != "true" {
			fmt.Fprintf(w, "%s\t%s\t%s\n",
				rec[0],
				rec[1],
				rec[2],
			)
		}
	}
}

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
			w := tabwriter.NewWriter(os.Stdout, 0, 4, 5, ' ', 0)
			defer w.Flush()
			printRecords(w, records, all)
		}

	},
}

func init() {
	listCommand.Flags().BoolVarP(&all, "all", "a", false, "Shows all flag task items (alias: -a)")
	rootCmd.AddCommand(listCommand)
}
