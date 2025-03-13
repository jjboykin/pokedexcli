package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jjboykin/pokedexcli/internal/pokeapi"
	"github.com/jjboykin/pokedexcli/internal/pokecache"
)

var commandRegistry map[string]cliCommand
var pokedex = make(map[string]pokeapi.Pokemon, 0)

func main() {

	commandRegistry = map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "List of all the Pokémon located at a given location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Prints the name, height, weight, stats and type(s) of a given Pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of all the Pokemon you have caught",
			callback:    commandPokedex,
		},
	}

	pokeCache := pokecache.NewCache(5 * time.Minute)

	locationConfigPtr := &config{
		Next:  "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Cache: &pokeCache,
	}

	//Wait for user input using bufio.NewScanner
	scanner := bufio.NewScanner(os.Stdin)

	// This loop will execute once for every command the user types in (we don't want to exit the program after just one command)
	for {

		fmt.Print("Pokedex > ")

		// Use the scanner's .Scan and .Text methods to get the user's input as a string
		scanner.Scan()
		input := scanner.Text()

		// Clean the users input by trimming any leading or trailing whitespace, and converting it to lowercase.
		words := cleanInput(input)

		// Capture the first "word" of the input and use it to print: Your command was: <first word>
		if len(words) > 0 {
			function, ok := commandRegistry[words[0]]
			if ok {
				var args []string
				if len(words) > 1 {
					args = words[1:]
				} else {
					args = make([]string, 0)
				}

				err := function.callback(locationConfigPtr, args)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Printf("Unknown command: %s\n", words[0])
			}
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(configPtr *config, args []string) error
}

type config struct {
	Previous string
	Next     string
	Cache    *pokecache.Cache
}

func commandCatch(configPtr *config, args []string) error {
	if len(args) == 0 {
		return errors.New("catch command requires a pokemon <name> or <id> argument")
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]
	pokemonData, err := pokeapi.GetPokemon(url, configPtr.Cache)
	if err != nil {
		return errors.New("pokemon not found")
	}

	baseCatchPercentage := 80
	catchPercentage := baseCatchPercentage - int((float64(pokemonData.BaseExperience) * 0.3))

	fmt.Printf("Throwing a Pokeball at %s...", pokemonData.Name)
	randomNumber := rand.Intn(100)

	if randomNumber < catchPercentage {
		fmt.Printf("Success! %s was caught and added to your Pokedex.\n", pokemonData.Name)
		pokedex[pokemonData.Name] = pokemonData
		fmt.Println("You may now view it with the 'inspect' command.")
	} else {
		fmt.Printf("Oh, no! %s escaped!\n", pokemonData.Name)
	}

	return nil
}

func commandExit(configPtr *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("error exiting program")
}

func commandExplore(configPtr *config, args []string) error {
	if len(args) == 0 {
		return errors.New("explore command requires a <name> or <id> argument")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	locationAreaData, err := pokeapi.GetLocationArea(url, configPtr.Cache)
	if err != nil {
		return errors.New("area not found")
	}

	fmt.Printf("Exploring %s [area %d]...\n", locationAreaData.Name, locationAreaData.ID)

	if len(locationAreaData.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
	} else {
		fmt.Println("No Pokemon found in the area")
	}
	// Parse the Pokemon's names from the response and display them to the user.
	for _, encounter := range locationAreaData.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

func commandHelp(configPtr *config, args []string) error {
	fmt.Println("----------------------------------------")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()

	if !(len(commandRegistry) > 0) {
		return errors.New("no valid command registry")
	}

	keys := make([]string, 0, len(commandRegistry))
	for k := range commandRegistry {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		command := commandRegistry[k]
		var name string
		if len(command.name) > 3 {
			name = "  %s: \t%s\n"
		} else {
			name = "  %s: \t\t%s\n"
		}
		fmt.Printf(name, command.name, command.description)
	}
	fmt.Println()

	return nil
}

func commandInspect(configPtr *config, args []string) error {

	if len(args) == 0 {
		return errors.New("inspect command requires a pokemon <name> argument")
	}

	pokemonData, ok := pokedex[args[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("----------------------------------------")

	idString := `Name: %s

Dimensions:
  -height: %d
  -weight: %d`

	fmt.Printf(idString+"\n", pokemonData.Name, pokemonData.Height, pokemonData.Weight)
	fmt.Println()
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println()
	fmt.Println("Types:")
	for _, attr := range pokemonData.Types {
		fmt.Printf("  -%s \n", attr.Type.Name)
	}

	fmt.Println()
	return nil
}

func commandMap(configPtr *config, args []string) error {
	url := configPtr.Next
	locationAreasData, err := pokeapi.GetLocationAreas(url, configPtr.Cache)
	if err != nil {
		return errors.New("error loading map page")
	}

	// Update the next and previous URLs in the config
	if locationAreasData.Next != nil {
		configPtr.Next = *locationAreasData.Next // Dereference to get the string value
	} else {
		configPtr.Next = "" // Or handle the nil case appropriately
	}
	if locationAreasData.Previous != nil {
		configPtr.Previous = *locationAreasData.Previous // Dereference to get the string value
	} else {
		configPtr.Previous = "" // Or handle the nil case appropriately
	}

	// Print the location names
	for _, location := range locationAreasData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(configPtr *config, args []string) error {
	if configPtr.Previous == "" {
		fmt.Println("You're already on the first page!")
		return nil
	}

	url := configPtr.Previous
	locationAreasData, err := pokeapi.GetLocationAreas(url, configPtr.Cache)
	if err != nil {
		return errors.New("error loading map page")
	}

	// Update the next and previous URLs in the config
	if locationAreasData.Next != nil {
		configPtr.Next = *locationAreasData.Next // Dereference to get the string value
	} else {
		configPtr.Next = "" // Or handle the nil case appropriately
	}
	if locationAreasData.Previous != nil {
		configPtr.Previous = *locationAreasData.Previous // Dereference to get the string value
	} else {
		configPtr.Previous = "" // Or handle the nil case appropriately
	}

	// Print the location names
	for _, location := range locationAreasData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandPokedex(configPtr *config, args []string) error {

	if len(pokedex) == 0 {
		fmt.Println("You have no Pokemon in your Pokedex.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Printf("  -%s\n", pokemon.Name)
	}

	return nil
}

func cleanInput(text string) []string {
	var words []string
	for _, word := range strings.Fields(text) {
		words = append(words, strings.Trim(strings.ToLower(word), " "))
	}
	return words
}
