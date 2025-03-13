
    - You'll need to use the PokeAPI (https://pokeapi.co/) location-area endpoint (https://pokeapi.co/docs/v2#location-areas) to get the location areas. Note that this is a different endpoint than the "location" endpoint. Calling the endpoint without an id will return a batch of location areas.

    - JSON lint is a useful tool for debugging JSON, it makes it easier to read.

    - JSON to Go (https://mholt.github.io/json-to-go/) a useful tool for converting JSON to Go structs. You can use it to generate the structs you'll need to parse the PokeAPI response. Keep in mind it sometimes can't know the exact type of a field that you want, because there are multiple valid options. For nullable strings, use *string.

    - Make basic API call function as a generic that the specific functions call

    - order pokedex by name or id, displaying what it is ordered by
    - add argument to specify sorting
    - format pokedex output like inspect
    - add a save file for pokedex
    - add different balls to increase catch chance
    - track current map zone beyond explore command output
    - limit catching of pokemon to those in current zone
    - query a list of pokemon by type
    - query a list of zones a given pokemon can be found in
    - query zones by region
    - track current region
    - list only zones in the current region by default in the map command
    - format map command output
    - add argument to map command list regions

    - Update the CLI to support the "up" arrow to cycle through previous commands
    - Simulate battles between pokemon
    - Add more unit tests
    - Refactor your code to organize it better and make it more testable
    - Keep pokemon in a "party" and allow them to level up
    - Allow for pokemon that are caught to evolve after a set amount of time
    - Persist a user's Pokedex to disk so they can save progress between sessions
    - Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
    - Random encounters with wild pokemon
    - Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon

    setup jira in wsl?