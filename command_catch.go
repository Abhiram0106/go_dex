package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(ctrl *Controller, parameters ...string) error {

	if len(parameters) != 1 {
		return errors.New("Enter a pokemon's name to catch")
	}

	res, err := ctrl.httpClient.CatchPokemon(parameters[0])

	if err != nil {
		return err
	}

	ctrl.pokedex[res.Name] = res

	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)

	luck := rand.Uint64()

	if res.BaseExperience < luck {
		fmt.Printf("%s was caught!\n", res.Name)
	} else {
		fmt.Printf("%s escaped!\n", res.Name)
	}

	return nil
}
