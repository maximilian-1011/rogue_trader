// Package creation handles all functions relating to characer creation
package creation

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Character struct {
	rank             int
	attributes       map[string]int
	healthPoints     int
	insanityPoints   int
	corruptionPoints int
	fatePoints       int
	name             string
	careerPath       string
	talents          []string
	skillsBasic      []string
	skillsTrained    []string
	skills10         []string
	skills20         []string
	equipment        []string
}

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

func NewCharacter() Character {
	return Character{
		rank:             1,
		attributes:       map[string]int{},
		healthPoints:     0,
		insanityPoints:   0,
		corruptionPoints: 0,
		fatePoints:       0,
		name:             "",
		careerPath:       "",
		talents:          []string{},
		skillsBasic:      []string{},
		skillsTrained:    []string{},
		skills10:         []string{},
		skills20:         []string{},
		equipment:        []string{},
	}
}

func CreateCharacter(char *Character) *Character {
	char.attributes = getAttributes()
	fmt.Println("Creating Character...")
	attributes := attributeList()
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter character name: ")
	scanner.Scan()
	char.name = scanner.Text()
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
				return char
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
	return char
}

func PresentCharacter(char *Character) {
	fmt.Println()
	fmt.Printf("Name: %s\n", char.name)
	fmt.Printf("Career Path: %s\n", char.careerPath)
	fmt.Printf("Rank: %v\n", char.rank)
	fmt.Printf("Health: %d\n", char.healthPoints)
	fmt.Printf("Insanity Points: %d\n", char.insanityPoints)
	fmt.Printf("Corruption Points: %d\n", char.corruptionPoints)
	fmt.Printf("Fate Points: %d\n", char.fatePoints)
	fmt.Println()
	fmt.Println("Attributes")
	for _, attribute := range attributeList() {
		fmt.Printf("%s: %d\n", attribute, char.attributes[attribute])
	}
	fmt.Println()
	fmt.Println("Talents")
	for _, talent := range char.talents {
		fmt.Println(talent)
	}
	fmt.Println()
	fmt.Println("Basic Skills")
	for _, basicSkill := range char.skillsBasic {
		fmt.Println(basicSkill)
	}
	fmt.Println()
	fmt.Println("Trained Skills")
	for _, trainedSkill := range char.skillsTrained {
		fmt.Println(trainedSkill)
	}
	fmt.Println()
	fmt.Println("Skills +10")
	for _, skill10 := range char.skills10 {
		fmt.Println(skill10)
	}
	fmt.Println()
	fmt.Println("Skills +20")
	for _, skill20 := range char.skills20 {
		fmt.Println(skill20)
	}
	fmt.Println()
	fmt.Println("Equipment")
	for _, item := range char.equipment {
		fmt.Println(item)
	}
	fmt.Println()
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Split(output, " ")
	return words
}
