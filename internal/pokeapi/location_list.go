package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokeApiResponse, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := PokeApiResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokeApiResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiResponse{}, err
	}

	res, errRes := c.httpClient.Do(req)
	if errRes != nil {
		return PokeApiResponse{}, errRes
	}
	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)
	if readError != nil {
		return PokeApiResponse{}, readError
	}

	locations := PokeApiResponse{}
	if err := json.Unmarshal(body, &locations); err != nil {
		return PokeApiResponse{}, err
	}

	c.cache.Add(url, body)
	return locations, nil
}
