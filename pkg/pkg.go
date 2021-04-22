package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetSingleLineInput(scn *bufio.Scanner) string {
	scn.Scan()
	return scn.Text()
}

// GetMultiLineInput will return the multiline input
// Input can be terminated with a Ctrl-D (EOF) on an empty line
// **OR**
// an empty line terminated with a ASCII Group Separator Ctrl-] and newline
func GetMultiLineInput(scn *bufio.Scanner) string {
	var lines []string
	terminator := false

	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			// ASCII Group Separator (GS ^]): ctrl-]
			if strings.HasSuffix(line, "\x1D") || strings.HasSuffix(line, "\x04") {
				trimLine := strings.TrimSuffix(line, "\x1D")
				if !(trimLine == "") {
					lines = append(lines, trimLine)
				}
				terminator = true
				break
			}
		}
		lines = append(lines, line)
	}

	if err := scn.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return strings.Join(lines, "\n")
	}
	// If we didn't break out of the scan loop
	// due to a Group Separator input character on an empty line
	// then we need to append the last line that had the EOF (Ctrl-D) character
	if !terminator {
		lines = append(lines, scn.Text())
	}

	return strings.Join(lines, "\n")
}

// flags but no args
func GetFlags() []string {
	var args []string
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-") {
			args = append(args, arg)
		}
	}
	return args
}
