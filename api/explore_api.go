package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreAPI(name *string) (LocationAreaDetails, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + *name

	if data, ok := c.cache.Get(url); ok {
		var locationAreaDetails LocationAreaDetails

		if err := json.Unmarshal(data, &locationAreaDetails); err != nil {
			return LocationAreaDetails{}, err
		}

		// fmt.Println("Explore cache used....")

		return locationAreaDetails, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetails{}, err
	}
	defer res.Body.Close()

	var locationAreaDetails LocationAreaDetails

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	c.cache.Add(url, data)
	// fmt.Println("Added Explore data to the cash")

	if err := json.Unmarshal(data, &locationAreaDetails); err != nil {
		return LocationAreaDetails{}, err
	}

	return locationAreaDetails, nil
}
