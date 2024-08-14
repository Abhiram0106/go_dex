package main

import (
	"github.com/Abhiram0106/go_dex/internal/pokeapi"
	"time"
)

func main() {
	ctrl := &Controller{
		httpClient:  pokeapi.NewClient(10*time.Second, 5*time.Second),
		previousURL: nil,
		nextURL:     nil,
		pokedex:     make(map[string]pokeapi.CatchPokemonResponse),
	}
	StartRepl(ctrl)
}
