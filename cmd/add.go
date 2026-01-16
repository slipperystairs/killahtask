package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// func loadFile(filepath string) (*os.File, error) {
// 	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open file for reading")
// 	}

//     // Exclusive lock obtained on the file descriptor
// 	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
// 		_ = f.Close()
// 		return nil, err
// 	}

// 	return f, nil
// }

// func closeFile(f *os.File) error {
// 	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
// 	return f.Close()
// }

func HasUniqueDescription(task string, records [][]string) bool {
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

func GetNow() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// todo => Locking the file for writes, to avoid race conditions
var addCommand = &cobra.Command{
	Use:     "add",
	Short:   "Adds a new item",
	Aliases: []string{"a"},
	Long:    `This command will add a item to your list`,
	Run: func(cmd *cobra.Command, args []string) {
		// now := time.Now()
		// nowStr := now.Format(time.RFC3339)
		// parsed, _ := time.Parse(time.RFC3339, nowStr)
		// fmt.Printf("now %v\n", now)
		// fmt.Printf("nowStr %s\n", nowStr)
		// fmt.Printf("parsed %v\n", parsed)
		// timeDiff := timediff.TimeDiff(parsed)
		// fmt.Printf("difference: %v\n", timeDiff)
		currUsr, err := user.Current()
		checkError(err)
		var earlyExit bool = false

		if len(args) == 0 {
			PrintUsageMsg("add", "add_none")
			earlyExit = true
		} else if len(args) > 1 {
			PrintUsageMsg("add", "add_to_many")
			earlyExit = true
		}

		if earlyExit {
			os.Exit(1)
		}

		var fileExist bool = true
		var fileName string = "killahtask_" + currUsr.Username + ".csv"
		var filePath string = filepath.Join(currUsr.HomeDir, fileName)
		_, err = os.Stat(filePath)
		if err != nil {
			fileExist = false
		}

		if !fileExist {
			file, err := os.Create(filePath)
			checkError(err)
			defer file.Close()

			records := [][]string{
				{"task_id", "description", "created", "completed"},
				{"0", args[0], GetNow(), "false"},
			}
			w := csv.NewWriter(file)
			w.WriteAll(records)
			checkError(w.Error())
		} else {
			file, err := os.Open(filePath)
			checkError(err)

			csvReader := csv.NewReader(file)
			// Read all the records from the CSV file
			records, err := csvReader.ReadAll()
			file.Close()
			checkError(err)

			if len(records) > 0 {
				lastId, err := strconv.Atoi(records[len(records)-1][0])
				checkError(err)
				newId := strconv.Itoa(lastId + 1)

				if !HasUniqueDescription(args[0], records) {
					fmt.Printf("Description isn't unique! \"%s\" already exist!\n", args[0])
					os.Exit(1)
				}

				records = append(records, []string{newId, args[0], GetNow(), "false"})
				// fmt.Printf("new records %v\n", records)
				file, err := os.Create(filePath)
				checkError(err)
				defer file.Close()

				csvWriter := csv.NewWriter(file)
				csvWriter.WriteAll(records)
				checkError(csvWriter.Error())
			}
		}
		//  To see your task run: \"killahtask list\" to see your task.
		fmt.Println("Record added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}
