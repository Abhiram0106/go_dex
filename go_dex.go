package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Abhiram0106/go_dex/internal/pokeapi"
)

type Command struct {
	name        string
	description string
	command     func(ctrl *Controller, parameter *[]string) error
}

type Controller struct {
	httpClient  pokeapi.Client
	previousURL *string
	nextURL     *string
}

func StartRepl(ctrl *Controller) {

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("go_dex > ")

		reader.Scan()
		capturedText := reader.Text()
		if len(capturedText) == 0 {
			continue
		}

		input := cleanInput(&capturedText)
		command, exists := getCommands()[(*input)[0]]
		if !exists {
			fmt.Println("Unkown Command")
			continue
		}
		err := command.command(ctrl, input)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Prints all available commands",
			command:     commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit go_dex",
			command:     commandExit,
		},
		"map": {
			name:        "map",
			description: "List the next 20 locations",
			command:     commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 locations",
			command:     commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List the Pokemon of a given area. Ex: explore pastoria-city-area",
			command:     commandExplore,
		},
	}
}

func cleanInput(input *string) *[]string {
	output := strings.ToLower(*input)
	words := strings.Fields(output)
	return &words
}
