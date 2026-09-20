package creation

import (
	"fmt"
	"math/rand"
)

const (
	BALLISTIC    = "Ballistic Skill"
	WEAPON       = "Weapon Skill"
	STRENGTH     = "Strenght"
	TOUGHNESS    = "Toughness"
	AGILITY      = "Agility"
	PERCEPTION   = "Perception"
	INTELLIGENCE = "Intelligence"
	WILLPOWER    = "Willpower"
	FELLOWSHIP   = "Fellowship"
)

func attributeList() []string {
	return []string{
		BALLISTIC,
		WEAPON,
		STRENGTH,
		TOUGHNESS,
		AGILITY,
		PERCEPTION,
		INTELLIGENCE,
		WILLPOWER,
		FELLOWSHIP,
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

			choice := GetUserInput()[0]
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

	fmt.Println()

	if total >= 100 {
		for i := range attributes {
			fmt.Printf("%s: %d\n", attributes[i], rolls[i])
		}
		fmt.Printf("Your total is: %d\n", total)
		fmt.Println()
	}
	for i, attribute := range attributes {
		char.Attributes[attribute] += rolls[i]
	}

	availableSwaps := 2
	for availableSwaps > 0 {
		fmt.Printf("You have %d swaps remaining.", availableSwaps)
		fmt.Print("Would you like to swap? (yes/no): ")
		input := GetUserInput()[0]

		switch input {
		case "yes":
			stat1 := querryAttribute()
			stat2 := querryAttribute()
			swapStats(char, stat1, stat2)
			availableSwaps -= 1
		case "no":
			availableSwaps = 0
		}
	}
}

func swapStats(char *Character, stat1, stat2 string) {
	char.Attributes[stat1], char.Attributes[stat2] = char.Attributes[stat2], char.Attributes[stat1]
}
