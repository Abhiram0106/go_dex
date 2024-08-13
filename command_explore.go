package main

import (
	"errors"
	"fmt"
)

func commandExplore(ctrl *Controller, input ...string) error {

	if len(input) != 1 {
		return errors.New("Please specify a location. Type \"help\" for more info")
	}

	location := input[0]

	res, err := ctrl.httpClient.GetPokemonAtLocation(location)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Pokemon found:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
