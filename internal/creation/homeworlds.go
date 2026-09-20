package creation

import "fmt"

type homeworld struct {
	name     string
	callback func(*Character) error
}

func applyHomeworld(char *Character) error {
	homeworlds := getHomeworlds()
	fmt.Println("Pick your Home World:")
	for name := range homeworlds {
		fmt.Print(name + " ")
	}
	fmt.Println()
	for {
		fmt.Print("Home World > ")
		input := GetName()

		home, exists := homeworlds[input]
		if exists {
			err := home.callback(char)
			if err != nil {
				return err
			}
			return nil
		} else {
			fmt.Println()
			fmt.Println("Unknown home world, please copy the given name case sensitive")
			fmt.Println()
		}
	}
}

func getHomeworlds() map[string]homeworld {
	return map[string]homeworld{
		"Death World": {
			name:     "Death World",
			callback: applyDeathWorld,
		},
	}
}

func applyDeathWorld(char *Character) error {
	char.Attributes[STRENGTH] += 5
	char.Attributes[TOUGHNESS] += 5
	char.Attributes[WILLPOWER] -= 5
	char.Attributes[FELLOWSHIP] -= 5
	err := skill{
		name:      "Survival",
		attribute: INTELLIGENCE,
		level:     TRAINED,
	}.apply(char)
	if err != nil {
		return err
	}

	pickChoice(char, map[string]choice{
		"Jaded": talent{
			name: "Jaded",
		},
		"Resistance (Poisons)": talent{
			name: "Resistance (Poisons)",
		},
	})
	err = talent{
		name: "Melee Weapon Training (Primitive)",
	}.apply(char)
	if err != nil {
		return err
	}

	err = talent{
		name: "Paranoid",
	}.apply(char)
	if err != nil {
		return err
	}

	err = talent{
		name: "Survivor",
	}.apply(char)
	if err != nil {
		return err
	}

	char.HealthPoints = roll(1, 5) + 2 + (char.Attributes[TOUGHNESS] / 10)
	fate := roll(1, 10)
	if fate <= 5 {
		char.FatePoints = 2
	} else {
		char.FatePoints = 3
	}

	return nil
}
