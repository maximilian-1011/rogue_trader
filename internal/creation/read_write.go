package creation

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func CharacterWrite(char *Character, path string) error {
	words := CleanInput(char.Name)
	fileName := strings.Join(words, "_") + ".json"
	fullPath := path + "/" + fileName
	characterData, err := export(char)
	if err != nil {
		return err
	}
	fmt.Printf("Saving %s in %s\n", char.Name, fullPath)
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Println("A character of that name already exists.")
		fmt.Print("Do you wish to overwrite? (yes/no) ")
		choice := GetUserInput()[0]
		switch choice {
		case "yes":
			err := os.WriteFile(fullPath, characterData, 0o644)
			if err != nil {
				return err
			}
		case "no":
			fmt.Println("Character was not written!")
			return nil
		}
	} else if errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(fullPath, characterData, 0o644)
		if err != nil {
			return err
		}
		fmt.Println("Character was written succesfully!")
	}
	return nil
}

func CharacterRead(char *Character, path string) error {
	fmt.Print("Please enter the character's name: ")
	words := GetUserInput()
	fileName := strings.Join(words, "_") + ".json"
	fullPath := path + "/" + fileName
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, char)
	if err != nil {
		return err
	}

	return nil
}
