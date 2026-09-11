package creation

type motivation struct {
	name     string
	choices  [][]any
	callback func(*Character, [][]any) error
}

func getMotivation() map[string]motivation {
	return map[string]motivation{
		"endurance": {
			name:     "Endurance",
			choices:  [][]any{},
			callback: applyEndurance,
		},
		"fortune": {
			name:     "Fortune",
			choices:  [][]any{},
			callback: applyFortune,
		},
		"vengeance": {
			name:     "Vengeance",
			choices:  [][]any{},
			callback: applyVengeance,
		},
		"renown": {
			name:     "Renown",
			choices:  [][]any{},
			callback: applyRenown,
		},
		"pride": {
			name:     "Pride",
			choices:  [][]any{},
			callback: applyPride,
		},
		"prestige": {
			name:     "Prestige",
			choices:  [][]any{},
			callback: applyPrestige,
		},
	}
}

func applyEndurance(char *Character, choices [][]any) error {
	return nil
}

func applyFortune(char *Character, choices [][]any) error {
	return nil
}

func applyVengeance(char *Character, choices [][]any) error {
	return nil
}

func applyRenown(char *Character, choices [][]any) error {
	return nil
}

func applyPride(char *Character, choices [][]any) error {
	return nil
}

func applyPrestige(char *Character, choices [][]any) error {
	return nil
}
