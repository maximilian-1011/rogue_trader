package creation

type birthright struct {
	name     string
	choices  [][]any
	callback func(*Character, [][]any) error
}

func getBirthrights() map[string]birthright {
	return map[string]birthright{
		"scavenger": {
			name:     "Scavenger",
			choices:  [][]any{},
			callback: applyScavenger,
		},
		"scapegrace": {
			name:     "Scapegrace",
			choices:  [][]any{},
			callback: applyScapegrace,
		},
		"stubjack": {
			name:     "Stubjack",
			choices:  [][]any{},
			callback: applyStubjack,
		},
		"child": {
			name:     "Child of the Creed",
			choices:  [][]any{},
			callback: applyChildOfTheCreed,
		},
		"savant": {
			name:     "Savant",
			choices:  [][]any{},
			callback: applySavant,
		},
		"vaunted": {
			name:     "Vaunted",
			choices:  [][]any{},
			callback: applyVaunted,
		},
	}
}

func applyScavenger(char *Character, choices [][]any) error {
	return nil
}

func applyScapegrace(char *Character, choices [][]any) error {
	return nil
}

func applyStubjack(char *Character, choices [][]any) error {
	return nil
}

func applyChildOfTheCreed(char *Character, choices [][]any) error {
	return nil
}

func applySavant(char *Character, choices [][]any) error {
	return nil
}

func applyVaunted(char *Character, choices [][]any) error {
	return nil
}
