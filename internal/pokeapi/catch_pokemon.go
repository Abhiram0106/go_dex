package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) CatchPokemon(nameOfPokemon string) (CatchPokemonResponse, error) {

	url := baseUrl + pokemon + "/" + nameOfPokemon

	res, err := c.httpClient.Get(url)

	if err != nil {
		return CatchPokemonResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return CatchPokemonResponse{}, err
	}

	catchPoke := CatchPokemonResponse{}

	unMarshallErr := json.Unmarshal(body, &catchPoke)

	if unMarshallErr != nil {
		return CatchPokemonResponse{}, unMarshallErr
	}

	return catchPoke, nil
}
