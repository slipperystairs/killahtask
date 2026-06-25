package chodesay

import (
	"fmt"
	"slices"
	"strings"
)

func ChodeSay(lines []string) {
	// Find the longest length in our slice of strings
	longestLine := slices.MaxFunc(lines, func(a string, b string) int {
		return len(a) - len(b)
	})
	// 2 accounts for the extra space an pipe
	boxWidth := len(longestLine)
	if boxWidth % 2 == 0 {
		boxWidth += 1
	}
	repeated := strings.Repeat("-", boxWidth/2 + 1) + ";" + strings.Repeat("-", boxWidth/2 + 1)
	fmt.Printf(" /%s\\\n", repeated)
	repeated2 := strings.Repeat(" ", boxWidth/2 + 1) + " " + strings.Repeat(" ", boxWidth/2 + 1)
	fmt.Printf("/ %s \\\n", repeated2)
	for _, line := range lines {
		padLen := boxWidth - len(line)
		if padLen < 0 {
			padLen = 0
		}

		padding := strings.Repeat(" ", padLen)
		str := " | " + line + padding + " |\n"
		fmt.Printf("%s", str)
	}
	repeated3 := strings.Repeat(" ", boxWidth/2) + ") (" + strings.Repeat(" ", boxWidth/2)
	fmt.Printf("( %s )\n", repeated3)

}
