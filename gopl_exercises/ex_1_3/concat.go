package concat

import "strings"

func concatenatingConcat(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	return s
}

func joiningConcat(args []string) string {
	return strings.Join(args[1:], " ")
}
