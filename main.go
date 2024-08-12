package main

import (
	"github.com/Abhiram0106/go_dex/internal/pokeapi"
	"time"
)

func main() {
	ctrl := &Controller{
		httpClient:  pokeapi.NewClient(10 * time.Second),
		previousURL: nil,
		nextURL:     nil,
	}
	StartRepl(ctrl)
}
