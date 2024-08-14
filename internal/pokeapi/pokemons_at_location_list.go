package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) GetPokemonAtLocation(location_name string) (PokemonAtLocationResponse, error) {

	url := baseUrl + locations + "/" + location_name

	cache, exists := c.cache.Get(location_name)
	if exists {
		pokemon_at_loc := PokemonAtLocationResponse{}
		marshalError := json.Unmarshal(cache, &pokemon_at_loc)

		if marshalError != nil {
			return PokemonAtLocationResponse{}, marshalError
		}

		fmt.Println("CACHE HIT")

		return pokemon_at_loc, nil
	}

	res, err := c.httpClient.Get(url)

	if err != nil {
		return PokemonAtLocationResponse{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonAtLocationResponse{}, err
	}

	if res.StatusCode > 299 {
		return PokemonAtLocationResponse{}, errors.New(fmt.Sprintf("Get Pokemons at location request failed with status code %d and \nbody: %s", res.StatusCode, body))
	}

	pokemon_at_loc := PokemonAtLocationResponse{}
	marshalError := json.Unmarshal(body, &pokemon_at_loc)

	if marshalError != nil {
		return PokemonAtLocationResponse{}, marshalError
	}

	c.cache.Add(location_name, body)

	return pokemon_at_loc, nil
}
