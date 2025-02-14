package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationArea(arg *string) (LocationAreaStruct, error) {
	url := baseURL + "/location-area/" + *arg

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreaStruct{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreaStruct{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaStruct{}, err
	}

	res, errRes := c.httpClient.Do(req)
	if errRes != nil {
		return LocationAreaStruct{}, errRes
	}
	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)
	if readError != nil {
		return LocationAreaStruct{}, readError
	}

	locations := LocationAreaStruct{}
	if err := json.Unmarshal(body, &locations); err != nil {
		return LocationAreaStruct{}, err
	}

	c.cache.Add(url, body)
	return locations, nil
}
