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

func GetMultiLineInput(scn *bufio.Scanner) []string {
	var lines []string

	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			// Group Separator (GS ^]): ctrl-]
			if strings.HasSuffix(line, "\x1D") || strings.HasSuffix(line, "\x04") {
				trimLine := strings.TrimSuffix(line, "\x1D")
				if !(trimLine == "") {
					lines = append(lines, trimLine)
				}
				break
			}
		}
		lines = append(lines, line)
	}

	if err := scn.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return lines
	}
	if len(lines) == 0 {
		return lines
	}

	return lines
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
