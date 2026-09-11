package creation

import (
	"strings"
)

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Split(output, " ")
	return words
}
