package creation

type trial struct {
	name     string
	choices  [][]any
	callback func(*Character, [][]any) error
}

func getTrialsAndTravails() map[string]trial {
	return map[string]trial{
		"the": {
			name:     "The Hand of War",
			choices:  [][]any{},
			callback: applyTheHandOfWar,
		},
		"press": {
			name:     "Press Ganged",
			choices:  [][]any{},
			callback: applyPressGanged,
		},
		"calamity": {
			name:     "Calamity",
			choices:  [][]any{},
			callback: applyCalamity,
		},
		"ship": {
			name:     "Ship Lorn",
			choices:  [][]any{},
			callback: applyShipLorn,
		},
		"dark": {
			name:     "Dark Voyage",
			choices:  [][]any{},
			callback: applyDarkVoyage,
		},
		"high": {
			name:     "High Vendetta",
			choices:  [][]any{},
			callback: applyHighVendetta,
		},
	}
}

func applyTheHandOfWar(char *Character, choices [][]any) error {
	return nil
}

func applyPressGanged(char *Character, choices [][]any) error {
	return nil
}

func applyCalamity(char *Character, choices [][]any) error {
	return nil
}

func applyShipLorn(char *Character, choices [][]any) error {
	return nil
}

func applyDarkVoyage(char *Character, choices [][]any) error {
	return nil
}

func applyHighVendetta(char *Character, choices [][]any) error {
	return nil
}
