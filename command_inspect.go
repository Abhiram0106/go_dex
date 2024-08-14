package main

import (
	"errors"
	"fmt"
)

func commandInspect(ctrl *Controller, parameters ...string) error {

	if len(parameters) != 1 {
		return errors.New("Enter the name of a pokemon")
	}

	pokemonName := parameters[0]

	pokemon, inPokedex := ctrl.pokedex[pokemonName]

	if !inPokedex {
		return errors.New("You haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, poke_type := range pokemon.Types {
		fmt.Printf("	-%s\n", poke_type.Type.Name)
	}
	return nil
}
