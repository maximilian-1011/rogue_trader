package main

import (
	"fmt"
	"os"
)

func commandExit(chr *Character) error {
	fmt.Println()
	fmt.Println("Closing Rogue Trader Cli")
	fmt.Println("Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}
