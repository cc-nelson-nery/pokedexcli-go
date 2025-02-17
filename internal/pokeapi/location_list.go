package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationList, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationList{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationList{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationList{}, err
	}

	res, errRes := c.httpClient.Do(req)
	if errRes != nil {
		return LocationList{}, errRes
	}
	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)
	if readError != nil {
		return LocationList{}, readError
	}

	locations := LocationList{}
	if err := json.Unmarshal(body, &locations); err != nil {
		return LocationList{}, err
	}

	c.cache.Add(url, body)
	return locations, nil
}
