package main

import (
	"errors"
	"fmt"
)

func commandExplore(ctrl *Controller, input *[]string) error {

	if len((*input)) < 2 {
		return errors.New("Please specify a location. Type \"help\" for more info")
	}

	location := (*input)[1]

	res, err := ctrl.httpClient.GetPokemonAtLocation(location)

	if err != nil {
		return err
	}

	for _, encounter := range res.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
