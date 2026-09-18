// Package creation handles all functions relating to characer creation
package creation

import (
	"encoding/json"
	"fmt"
)

type Character struct {
	Rank             int            `json:"rank"`
	Attributes       map[string]int `json:"attributes"`
	HealthPoints     int            `json:"health"`
	InsanityPoints   int            `json:"insanity"`
	CorruptionPoints int            `json:"corruption"`
	FatePoints       int            `json:"fate"`
	Name             string         `json:"name"`
	CareerPath       string         `json:"career"`
	Talents          []string       `json:"talents"`
	SkillsBasic      []string       `json:"basic"`
	SkillsTrained    []string       `json:"trained"`
	Skills10         []string       `json:"skills10"`
	Skills20         []string       `json:"skills20"`
	Equipment        []string       `json:"equipment"`
}

func NewCharacter() Character {
	return Character{
		Rank:             1,
		Attributes:       map[string]int{},
		HealthPoints:     0,
		InsanityPoints:   0,
		CorruptionPoints: 0,
		FatePoints:       0,
		Name:             "",
		CareerPath:       "",
		Talents:          []string{},
		SkillsBasic:      []string{},
		SkillsTrained:    []string{},
		Skills10:         []string{},
		Skills20:         []string{},
		Equipment:        []string{},
	}
}

func CreateCharacter(char *Character) *Character {
	char.Attributes = getAttributes()
	fmt.Println("Creating Character...")
	//attributes := attributeList()
	fmt.Print("Enter character name: ")
	char.Name = GetName()
	getStats(char)
	return char
}

func PresentCharacter(char *Character) {
	fmt.Println()
	fmt.Printf("Name: %s\n", char.Name)
	fmt.Printf("Career Path: %s\n", char.CareerPath)
	fmt.Printf("Rank: %v\n", char.Rank)
	fmt.Printf("Health: %d\n", char.HealthPoints)
	fmt.Printf("Insanity Points: %d\n", char.InsanityPoints)
	fmt.Printf("Corruption Points: %d\n", char.CorruptionPoints)
	fmt.Printf("Fate Points: %d\n", char.FatePoints)
	fmt.Println()
	fmt.Println("Attributes")
	for _, attribute := range attributeList() {
		fmt.Printf("%s: %d\n", attribute, char.Attributes[attribute])
	}
	fmt.Println()
	fmt.Println("Talents")
	for _, talent := range char.Talents {
		fmt.Println(talent)
	}
	fmt.Println()
	fmt.Println("Basic Skills")
	for _, basicSkill := range char.SkillsBasic {
		fmt.Println(basicSkill)
	}
	fmt.Println()
	fmt.Println("Trained Skills")
	for _, trainedSkill := range char.SkillsTrained {
		fmt.Println(trainedSkill)
	}
	fmt.Println()
	fmt.Println("Skills +10")
	for _, skill10 := range char.Skills10 {
		fmt.Println(skill10)
	}
	fmt.Println()
	fmt.Println("Skills +20")
	for _, skill20 := range char.Skills20 {
		fmt.Println(skill20)
	}
	fmt.Println()
	fmt.Println("Equipment")
	for _, item := range char.Equipment {
		fmt.Println(item)
	}
	fmt.Println()
}

func export(char *Character) ([]byte, error) {
	data, err := json.Marshal(char)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
