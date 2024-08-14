package main

import "fmt"

func commandPokedex(ctrl *Controller, parameters ...string) error {

	fmt.Println("Your Pokedex:")

	for _, pokemon := range ctrl.pokedex {
		fmt.Printf("	- %s\n", pokemon.Name)
	}

	return nil
}
