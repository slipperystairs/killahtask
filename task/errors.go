package task

import (
	"fmt"
	"os"
)

func MaybeDieAboutIt(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}
