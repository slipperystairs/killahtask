package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func writeCSV(file *os.File, records [][]string) {
	w := csv.NewWriter(file)
	w.WriteAll(records)
	checkError(w.Error())
}

func loadFile(filepath string) (*os.File, error) {
	// Open or create file if it doesn't exist.
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file for reading")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func uniqueDescription(task string, records [][]string) bool {
	var isUnique bool = true

	for _, record := range records {
		for _, r := range record {
			if task == r {
				isUnique = false
				break
			}
		}

		if !isUnique {
			break
		}
	}

	return isUnique
}

func Now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

var addCommand = &cobra.Command{
	Use:     "add",
	Short:   "Adds a new item",
	Aliases: []string{"a"},
	Long:    `This command will add a item to your list`,
	Run: func(cmd *cobra.Command, args []string) {
		currUsr, err := user.Current()
		checkError(err)
		var earlyExit bool = false

		if len(args) == 0 {
			PrintUsageMsg("add", "add_none")
			earlyExit = true
		} else if len(args) > 1 { // A shitty way of making the user wrap their command in double quotes lol
			PrintUsageMsg("add", "add_to_many")
			earlyExit = true
		}

		if earlyExit {
			os.Exit(1)
		}

		var description string = strings.TrimSpace(args[0])
		var fileExist bool = true
		var fileName string = "killahtask_" + currUsr.Username + ".csv"
		var filePath string = filepath.Join(currUsr.HomeDir, fileName)
		// os.OpenFile doesn't have a way of letting us know if the file already exist.
		_, err = os.Stat(filePath)
		if err != nil {
			fileExist = false
		}

		file, err := loadFile(filePath)
		// File will get closed even in the event of an error.
		defer closeFile(file)
		checkError(err)

		if !fileExist {
			records := [][]string{
				{"task_id", "description", "created", "completed"},
				{"0", description, Now(), "false"},
			}
			writeCSV(file, records)
		} else {

			csvReader := csv.NewReader(file)
			records, err := csvReader.ReadAll()
			checkError(err)

			if len(records) > 0 {
				// Get the last task_id used and increment it by one.
				lastId, err := strconv.Atoi(records[len(records)-1][0])
				checkError(err)
				newId := strconv.Itoa(lastId + 1)

				if !uniqueDescription(description, records) {
					fmt.Printf("Task description isn't unique! \"%s\" already exist.\n", description)
					os.Exit(1)
				}

				// Append the new record to the end of the slice.
				records = append(records, []string{newId, description, Now(), "false"})
				writeCSV(file, records)
			}
		}
		fmt.Printf("Task \"%s\" added successfully!\n", description)
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
