package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetLocations (pageURL *string) (AreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil{
		url = *pageURL
	}

	// New request
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaList{}, err
	}

	// Make request
	res, err := c.httpClient.Do(rq)
	if err != nil {
		return AreaList{}, err
	}
	
	// Decode the request
	list := AreaList{}
	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&list)
	if err != nil {
		return AreaList{}, err
	}

	return list, nil
}	