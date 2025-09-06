package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) Catch(pokemon_name *string) (*Pokemon, error) {
	url := baseURL + "/pokemon/"

	if pokemon_name == nil {
		return nil, errors.New("must have a valid name")
	}
	url = url + *pokemon_name
	
	pokemon := Pokemon{}
	if cached, ok := c.cache.Get(url); ok {
		json.Unmarshal(cached, &pokemon)
		return &pokemon, nil
	}
	
	// New request
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	// Make request
	res, err := c.httpClient.Do(rq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	//  Get, cache and return resutls
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	
	c.cache.Add(url, data)
	
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}