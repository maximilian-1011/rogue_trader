package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maximilian-1011/rogue_trader/internal/creation"
)

func startRepl() {
	chr := creation.NewCharacter()
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

		words := creation.CleanInput(scanner.Text())

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

type cliCommand struct {
	name        string
	description string
	callback    func(*creation.Character) error
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
