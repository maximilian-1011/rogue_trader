package main

import (
	"github.com/maximilian-1011/rogue_trader/internal/creation"
)

func commandBuild(char *creation.Character) error {
	creation.CreateCharacter(char)
	return nil
}

func commandDisplay(char *creation.Character) error {
	creation.PresentCharacter(char)
	return nil
}

func commandLoad(char *creation.Character) error {
	return nil
}

func commandSave(char *creation.Character) error {
	return nil
}
