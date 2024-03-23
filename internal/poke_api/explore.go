package poke_api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type LocationDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	VersionDetails []struct {
		EncounterDetails []struct {
			Chance          int   `json:"chance"`
			ConditionValues []any `json:"condition_values"`
			MaxLevel        int   `json:"max_level"`
			Method          struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"method"`
			MinLevel int `json:"min_level"`
		} `json:"encounter_details"`
		MaxChance int `json:"max_chance"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"version_details"`
}

func (c *Client) ExploreLocation(area string) (LocationDetails, error) {
	pageURL := baseURL + "/location-area/" + area
	data, exists := c.cache.Get(pageURL)
	if exists {
		var location LocationDetails
		err := json.Unmarshal(data, &location)
		if err != nil {
			log.Fatal(err)
			return LocationDetails{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest(http.MethodGet, pageURL, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return LocationDetails{}, err
	}
	c.cache.Set(pageURL, body)

	var location LocationDetails
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationDetails{}, err
	}

	c.cache.Set(pageURL, body)
	return location, nil
}
