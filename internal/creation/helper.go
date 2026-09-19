package creation

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Split(output, " ")
	return words
}

func GetUserInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := CleanInput(scanner.Text())
	return words
}

func GetName() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	return name
}

func roll(amount, faces int) int {
	return rand.Intn(faces-1) + 1
}
