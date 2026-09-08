package main

import "fmt"

func commandHelp(chr *Character) error {
	commands := getCommands()
	fmt.Println()
	fmt.Println("Welcome to Rogue Trader CLI")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
