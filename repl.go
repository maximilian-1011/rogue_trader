package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	chr := getCharacter()
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Rogue Trader Cli > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println(err)
			continue
		}

		words := cleanupInput(scanner.Text())

		commandName := words[0]

		command, exists := commands[commandName]
		if exists {
			err := command.callback(&chr)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println()
			fmt.Println("Unknown command! Type 'help' for list of available commands")
			fmt.Println()
			continue
		}
	}
}

func cleanupInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Split(output, " ")
	return words
}

type Character struct {
	name          string
	career        string
	rank          int
	health        int
	insanity      int
	corruption    int
	fate          int
	skillsBasic   []string
	skillsTrained []string
	skills10      []string
	skills20      []string
	xpSpent       int
	xpToSpend     int
}

func getCharacter() Character {
	return Character{}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Character) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Shows available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Closes Programm",
			callback:    commandExit,
		},
		"build": {
			name:        "build",
			description: "Begins Character creation",
			callback:    commandBuild,
		},
		"display": {
			name:        "display",
			description: "Displays current character",
			callback:    commandDisplay,
		},
		"save": {
			name:        "save",
			description: "Saves character as .json file",
			callback:    commandSave,
		},
		"load": {
			name:        "load",
			description: "Loads existing character",
			callback:    commandLoad,
		},
	}
}
