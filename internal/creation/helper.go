package creation

import (
	"bufio"
	"fmt"
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

func querryAttribute() string {
	for {
		fmt.Print("Please enter the shortened identifier: ")
		input := GetUserInput()[0]

		switch input {
		case "bs":
			return BALLISTIC
		case "ws":
			return WEAPON
		case "s":
			return STRENGTH
		case "t":
			return TOUGHNESS
		case "agi":
			return AGILITY
		case "per":
			return PERCEPTION
		case "int":
			return INTELLIGENCE
		case "wp":
			return WILLPOWER
		case "fel":
			return FELLOWSHIP
		default:
			fmt.Println(BALLISTIC + " (bs)")
			fmt.Println(WEAPON + " (ws)")
			fmt.Println(STRENGTH + " (s)")
			fmt.Println(TOUGHNESS + " (t)")
			fmt.Println(AGILITY + " (agi)")
			fmt.Println(PERCEPTION + " (per)")
			fmt.Println(INTELLIGENCE + " (int)")
			fmt.Println(WILLPOWER + " (wp)")
			fmt.Println(FELLOWSHIP + " (fel)")
			continue
		}
	}
}
