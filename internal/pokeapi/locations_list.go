package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) GetLocations(pagedUrl *string) (PokeLocationsResponse, error) {

	url := baseUrl + locations

	if pagedUrl != nil {
		url = *pagedUrl
	}

	cachedResult, exists := c.cache.Get(url)
	if exists {
		poke_locs := PokeLocationsResponse{}
		marshal_err := json.Unmarshal(cachedResult, &poke_locs)
		if marshal_err == nil {
			fmt.Println("CACHE HIT")
			return poke_locs, nil
		}
	}

	res, err := c.httpClient.Get(url)

	if err != nil {
		return PokeLocationsResponse{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return PokeLocationsResponse{}, err
	}

	if res.StatusCode > 299 {
		return PokeLocationsResponse{}, errors.New(fmt.Sprintln("Response failed with code %d and body\n%s", res.StatusCode, body))
	}

	poke_locs := PokeLocationsResponse{}
	marshal_err := json.Unmarshal(body, &poke_locs)

	if marshal_err != nil {
		return PokeLocationsResponse{}, marshal_err
	}

	c.cache.Add(url, body)

	return poke_locs, nil
}
