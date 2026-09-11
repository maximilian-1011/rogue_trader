package creation

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func attributeList() []string {
	return []string{
		"Ballistic Skill",
		"Weapon Skill",
		"Strength",
		"Toughness",
		"Agility",
		"Perception",
		"Intelligence",
		"Willpower",
		"Fellowship",
	}
}

func getAttributes() map[string]int {
	attributeMap := map[string]int{}
	attributes := attributeList()
	for _, attribute := range attributes {
		attributeMap[attribute] = 25
	}

	return attributeMap
}

func getStats(char *Character) {
	attributes := attributeList()
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	var rolls []int
	for total < 100 {
		total = 0
		rolls = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i := range attributes {
			rolls[i] = rand.Intn(20-2) + 2
			total += rolls[i]
		}
		exit := false
		if total < 100 {
			fmt.Println()
			for i := range attributes {
				fmt.Printf("%s: %d\n", attributes[i], rolls[i])
			}
			fmt.Printf("Your total is: %d\n", total)
			fmt.Print("Would you like to reroll? (yes/no): ")
			scanner.Scan()
			err := scanner.Err()
			if err != nil {
				fmt.Println(err)
				return
			}

			choice := CleanInput(scanner.Text())[0]
			switch choice {
			case "yes":
				continue
			case "no":
				exit = true
			default:
				continue
			}
		}
		fmt.Println()
		if exit {
			break
		}
	}
	if total >= 100 {
		for i := range attributes {
			fmt.Printf("%s: %d\n", attributes[i], rolls[i])
		}
		fmt.Printf("Your total is: %d\n", total)
		fmt.Println()
	}
	for i, attribute := range attributes {
		char.attributes[attribute] += rolls[i]
	}
}
