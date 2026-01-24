package task

import (
	"encoding/csv"
	"fmt"
	"os"
	"syscall"
)

func WriteCSV(file *os.File, records [][]string) error {
	// If the file already exist the file pointer is at byte 0.
	// Writing the file without truncating/seeking will duplicate header + rows.
	file.Truncate(0) // Cuts the file down to byte making an empty file.
	file.Seek(0, 0)  // Moves the file pointer back to the beginning

	w := csv.NewWriter(file)
	err := w.WriteAll(records)
	if err != nil {
		return err
	}

	return w.Error()
}

func LoadFile(filepath string) (*os.File, error) {
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

func CloseFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}
