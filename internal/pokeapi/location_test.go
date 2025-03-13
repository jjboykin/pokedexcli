package pokeapi

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jjboykin/pokedexcli/internal/pokecache"
)

func TestGetLocationAreas(t *testing.T) {
	cases := []struct {
		url      string
		expected string
	}{
		{
			url: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
			expected: `canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior`,
		},
		{
			url: "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
			expected: `mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
great-marsh-area-3
great-marsh-area-4
great-marsh-area-5
great-marsh-area-6
solaceon-ruins-2f
solaceon-ruins-1f
solaceon-ruins-b1f-a
solaceon-ruins-b1f-b
solaceon-ruins-b1f-c
solaceon-ruins-b2f-a
solaceon-ruins-b2f-b
solaceon-ruins-b2f-c
solaceon-ruins-b3f-a
solaceon-ruins-b3f-b
solaceon-ruins-b3f-c`,
		},
	}

	pokeCache := pokecache.NewCache(5 * time.Minute)

	for _, c := range cases {
		actual, err := GetLocationAreas(c.url, &pokeCache)
		if err != nil {
			if c.expected != "no location returned" {
				t.Errorf("no location returned")
			} else {
				continue
			}
		}
		actualString := ""

		// Print the location names
		for _, location := range actual.Results {
			actualString += location.Name + "\n"
		}
		fmt.Println(actualString)
		fmt.Println("----------------------------")
		fmt.Println(c.expected)
		if strings.Trim(actualString, "\n") != strings.Trim(c.expected, "\n") {
			t.Errorf("string mismatch")
		}

	}
}
