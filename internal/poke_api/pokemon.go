package poke_api

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	pageURL := fmt.Sprintf(baseURL+"/pokemon/%s", name)

	data, exists := c.cache.Get(pageURL)
	if exists {
		var pokemon Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest(http.MethodGet, pageURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Set(pageURL, body)

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

func (c *Client) CatchPokemon(name string) (Pokemon, bool, error) {
	pokemon, err := c.GetPokemon(name)
	if err != nil {
		return Pokemon{}, false, err
	}

	base_exp := pokemon.BaseExperience
	probability := 1 / (1 + math.Exp(float64(base_exp)/1000))

	random := rand.Float64()

	return pokemon, random <= probability, nil
}
