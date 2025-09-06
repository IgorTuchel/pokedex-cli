package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations (pageURL *string) (AreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil{
		url = *pageURL
	}

	locations := AreaList{}

	if cached, ok := c.cache.Get(url); ok {
		json.Unmarshal(cached, &locations)
		return locations, nil
	}

	// New request
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locations, err
	}

	// Make request
	res, err := c.httpClient.Do(rq)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()
	
	//  Get, cache and return resutls
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locations, err
	}
	c.cache.Add(url, data)
	
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return AreaList{}, err
	}

	return locations, nil
}	






