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
	SkillsBasic      []skill        `json:"basic"`
	SkillsTrained    []skill        `json:"trained"`
	Skills10         []skill        `json:"skills10"`
	Skills20         []skill        `json:"skills20"`
	Equipment        []string       `json:"equipment"`
}

func NewCharacter() Character {
	return Character{
		Rank:             1,
		Attributes:       getAttributes(),
		HealthPoints:     0,
		InsanityPoints:   0,
		CorruptionPoints: 0,
		FatePoints:       0,
		Name:             "",
		CareerPath:       "",
		Talents:          []string{},
		SkillsBasic:      []skill{},
		SkillsTrained:    []skill{},
		Skills10:         []skill{},
		Skills20:         []skill{},
		Equipment:        []string{},
	}
}

func clearCharacter(char *Character) {
	char.Rank = 1
	char.Attributes = getAttributes()
	char.HealthPoints = 0
	char.InsanityPoints = 0
	char.CorruptionPoints = 0
	char.FatePoints = 0
	char.Name = ""
	char.CareerPath = ""
	char.Talents = []string{}
	char.SkillsBasic = []skill{}
	char.SkillsTrained = []skill{}
	char.Skills10 = []skill{}
	char.Skills20 = []skill{}
	char.Equipment = []string{}
}

func CreateCharacter(char *Character) {
	clearCharacter(char)
	fmt.Println("Creating Character...")
	//attributes := attributeList()
	fmt.Print("Enter character name: ")
	char.Name = GetName()
	getStats(char)
	err := applyHomeworld(char)
	if err != nil {
		fmt.Printf("An error occured during home world application: %v", err)
	}
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
		fmt.Println(basicSkill.name)
	}
	fmt.Println()
	fmt.Println("Trained Skills")
	for _, trainedSkill := range char.SkillsTrained {
		fmt.Println(trainedSkill.name)
	}
	fmt.Println()
	fmt.Println("Skills +10")
	for _, skill10 := range char.Skills10 {
		fmt.Println(skill10.name)
	}
	fmt.Println()
	fmt.Println("Skills +20")
	for _, skill20 := range char.Skills20 {
		fmt.Println(skill20.name)
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
