package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
    Name string `json:"name"`
    Types []struct {
        Type struct {
            Name string `json:"name"`
        } `json:"type"`
    } `json:"types"`
}

func fetchPokemonData(id int) (*Pokemon, error) {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var pokemon Pokemon
    err = json.NewDecoder(resp.Body).Decode(&pokemon)
    if err != nil {
        return nil, err
    }

    return &pokemon, nil
}

func main() {
    // Example: Fetching data for the first 10 Pokémon
    for i := 1; i <= 10; i++ {
        pokemon, err := fetchPokemonData(i)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Println(pokemon.Name, pokemon.Types[0].Type.Name)
    }
}
