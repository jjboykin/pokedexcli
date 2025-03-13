
    - You'll need to use the PokeAPI (https://pokeapi.co/) location-area endpoint (https://pokeapi.co/docs/v2#location-areas) to get the location areas. Note that this is a different endpoint than the "location" endpoint. Calling the endpoint without an id will return a batch of location areas.

    - JSON lint is a useful tool for debugging JSON, it makes it easier to read.

    - JSON to Go (https://mholt.github.io/json-to-go/) a useful tool for converting JSON to Go structs. You can use it to generate the structs you'll need to parse the PokeAPI response. Keep in mind it sometimes can't know the exact type of a field that you want, because there are multiple valid options. For nullable strings, use *string.

    - Make basic API call function as a generic that the specific functions call

    - order pokedex by name or id, displaying what it is ordered by
    - add argument to specify sorting
    - format pokedex output like inspect

    - add a save file for your pokedex
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

    setup jira in wsl?