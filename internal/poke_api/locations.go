package poke_api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type LocationAreas struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) ListLocationAreas(url *string) (LocationAreas, error) {
	var pageURL string
	if url != nil {
		pageURL = *url
	} else {
		pageURL = baseURL + "/location-area"
	}

	data, exists := c.cache.Get(pageURL)
	if exists {
		var locationAreas LocationAreas
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			log.Fatal(err)
			return LocationAreas{}, err
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest(http.MethodGet, pageURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}
	c.cache.Set(pageURL, body)

	var locationAreas LocationAreas
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}

	return locationAreas, nil
}
