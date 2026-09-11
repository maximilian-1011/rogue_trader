package creation

type lureOfTheVoid struct {
	name     string
	choices  [][]any
	callback func(*Character, [][]any) error
}

func getLuresOfTheVoid() map[string]lureOfTheVoid {
	return map[string]lureOfTheVoid{
		"tainted": {
			name:     "Tainted",
			choices:  [][]any{},
			callback: applyTainted,
		},
		"criminal": {
			name:     "Criminal",
			choices:  [][]any{},
			callback: applyCriminal,
		},
		"renegade": {
			name:     "Renegade",
			choices:  [][]any{},
			callback: applyRenegade,
		},
		"duty": {
			name:     "Duty Bound",
			choices:  [][]any{},
			callback: applyDutyBound,
		},
		"zealot": {
			name:     "Zealot",
			choices:  [][]any{},
			callback: applyZealot,
		},
		"chosen": {
			name:     "Chosen by Destiny",
			choices:  [][]any{},
			callback: applyChosenByDestiny,
		},
	}
}

func applyTainted(char *Character, choices [][]any) error {
	return nil
}

func applyCriminal(char *Character, choices [][]any) error {
	return nil
}

func applyRenegade(char *Character, choices [][]any) error {
	return nil
}

func applyDutyBound(char *Character, choices [][]any) error {
	return nil
}

func applyZealot(char *Character, choices [][]any) error {
	return nil
}

func applyChosenByDestiny(char *Character, choices [][]any) error {
	return nil
}
