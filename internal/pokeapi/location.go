package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jjboykin/pokedexcli/internal/pokecache"
)

type Location struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
}

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
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
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(url string, cache *pokecache.Cache) (LocationArea, error) {
	var locationArea LocationArea

	// Try to get data from cache first
	if data, found := cache.Get(url); found {
		// Data found in cache, unmarshal it
		err := json.Unmarshal(data, &locationArea)
		if err == nil {
			// Successfully unmarshaled cached data
			//fmt.Println("Retrieving from cache...")
			return locationArea, nil
		}
		// If there was an error unmarshaling, we'll just fall through to making the request
	}

	// Cache miss or unmarshal error, make the API request
	//fmt.Printf("Requesting from %s...", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return locationArea, err
	}

	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return locationArea, err
	}

	// Decode the response
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return locationArea, err
	}

	// Add the response to the cache
	cache.Add(url, body)

	//fmt.Println("Data recieved...")
	return locationArea, nil

}

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreas, error) {
	var locationAreas LocationAreas

	// Try to get data from cache first
	if data, found := cache.Get(url); found {
		// Data found in cache, unmarshal it
		err := json.Unmarshal(data, &locationAreas)
		if err == nil {
			// Successfully unmarshaled cached data
			//fmt.Println("Retrieving from cache...")
			return locationAreas, nil
		}
		// If there was an error unmarshaling, we'll just fall through to making the request
	}

	// Cache miss or unmarshal error, make the API request
	//fmt.Printf("Requesting from %s...", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return locationAreas, err
	}

	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return locationAreas, err
	}

	// Decode the response
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return locationAreas, err
	}

	// Add the response to the cache
	cache.Add(url, body)

	//fmt.Println("Data recieved...")
	return locationAreas, nil

}
