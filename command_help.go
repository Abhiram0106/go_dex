package main

import "fmt"

func commandHelp(ctrl *Controller) error {
	fmt.Println("Welcome to go_dex!")
	fmt.Print("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Print(fmt.Sprintf("%s: %s\n", cmd.name, cmd.description))
	}
	fmt.Println()

	return nil
}
