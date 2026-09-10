package main

import (
	"fmt"
	"os"

	"github.com/maximilian-1011/rogue_trader/internal/creation"
)

func commandExit(chr *creation.Character) error {
	fmt.Println()
	fmt.Println("Closing Rogue Trader Cli")
	fmt.Println("Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}
