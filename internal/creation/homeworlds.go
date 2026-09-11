package creation

type homeworld struct {
	name     string
	choices  map[string][]string
	callback func(*Character, map[string][]string) error
}

func getHomeworlds() map[string]homeworld {
	return map[string]homeworld{
		"death world": {
			name:     "Death World",
			choices:  map[string][]string{},
			callback: applyDeathWorld,
		},
	}
}

func applyDeathWorld(char *Character, choices map[string][]string) error {
	return nil
}
