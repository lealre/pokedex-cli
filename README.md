# A Basic CLI in Go

This project consists of a basic CLI written in Go to display information about Pokémon using (guess what!?) the [PokéAPI](https://pokeapi.co/).

## Available Commands:

- **catch**: Try to catch a Pokémon.
- **inspect**: Check the stats of a Pokémon.
- **pokedex**: List all your caught Pokémon.
- **exit**: Exit the Pokédex.
- **help**: Display a help message.
- **map**: Display 20 locations.
- **mapb**: Display the previous 20 locations.
- **explore**: List all Pokémon located in a specific area.

The CLI uses an in-memory cache to store API response data, running in a goroutine that clears the cache every minute.

## Running the CLI:

```sh
go run .
```
