package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationAreas(cmmdURL *string) (LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if cmmdURL != nil {
		url = *cmmdURL
	}

	var (
		locations LocationArea
		data      []byte
	)

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}

		c.cache.Add(url, data)
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationArea{}, err
	}

	return locations, nil
}
