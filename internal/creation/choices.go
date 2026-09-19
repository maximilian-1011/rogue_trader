package creation

import (
	"fmt"
)

type choice interface {
	apply(*Character) error
}

type skill struct {
	name      string
	level     string
	attribute string
}

const (
	BASIC   = "basic"
	TRAINED = "trained"
	SKILL10 = "+10"
	SKILL20 = "+20"
)

func (s skill) apply(char *Character) error {
	switch s.level {
	case "basic":
		char.SkillsBasic = append(char.SkillsBasic, s)
	case "trained":
		char.SkillsTrained = append(char.SkillsTrained, s)
	case "+10":
		char.Skills10 = append(char.Skills10, s)
	case "+20":
		char.Skills20 = append(char.Skills20, s)
	default:
		fmt.Errorf("skill proficiency does not match the existing categories basic, trained, +10, or +20")
	}
	return nil
}

type talent struct {
	name string
}

func (t talent) apply(char *Character) error {
	char.Talents = append(char.Talents, t.name)
	return nil
}

type characteristicAdvance struct {
	name      string
	attribute string
	increase  int
}

func (c characteristicAdvance) apply(char *Character) error {
	if _, ok := char.Attributes[c.attribute]; ok {
		char.Attributes[c.attribute] += c.increase
		return nil
	} else {
		return fmt.Errorf("attribute does not match any known characteristic")
	}
}

type item struct {
	name string
}

func (i item) apply(char *Character) error {
	char.Equipment = append(char.Equipment, i.name)
	return nil
}

func pickChoice(char *Character, choices map[string]choice) {
	fmt.Println("Pick one of the following:")
	for name := range choices {
		fmt.Println(name)
	}
	for {
		fmt.Print("Choice > ")
		input := GetName()

		if c, ok := choices[input]; ok {
			c.apply(char)
			return
		} else {
			fmt.Println()
			fmt.Println("Please input the choice exactly as presented")
			fmt.Println()
		}
	}
}
