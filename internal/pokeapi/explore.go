package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) Explore(location *string) (Pokemons, error) {
	url := baseURL + "/location-area/"

	pokemons := Pokemons{}
	if location == nil {
		return pokemons, errors.New("invalid location")
	}
	url = url + *location

	if cached, ok := c.cache.Get(url); ok {
		json.Unmarshal(cached, &pokemons)
		return pokemons, nil
	}

	
	// New request
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemons, err
	}

	// Make request
	res, err := c.httpClient.Do(rq)
	if err != nil {
		return pokemons, err
	}
	defer res.Body.Close()
	
	//  Get, cache and return resutls
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemons, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return Pokemons{}, err
	}

	return pokemons, nil
}