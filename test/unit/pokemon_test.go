package pokemon_test

import (
	"testing"

	"MyGoProject/internal/pokemon"
)

// TestNewPokemon tests the NewPokemon function
func TestNewPokemon(t *testing.T) {
    id := 1
    name := "Bulbasaur"
    types := []pokemon.PokemonType{{Type: struct{ Name, URL string }{"Grass", "url"}}}

    pkmn := pokemon.NewPokemon(id, name, types)

    if pkmn.ID != id {
        t.Errorf("Expected ID %d; got %d", id, pkmn.ID)
    }

    if pkmn.Name != name {
        t.Errorf("Expected Name %s; got %s", name, pkmn.Name)
    }

    if len(pkmn.Types) != 1 || pkmn.Types[0].Type.Name != "Grass" {
        t.Errorf("Expected Types to contain 'Grass'; got %v", pkmn.Types)
    }
}

// TestPokemonTypeString tests the String method of the PokemonType struct
func TestPokemonTypeString(t *testing.T) {
    typeCase := pokemon.PokemonType{Type: struct{ Name, URL string }{"Fire", "url"}}

    expected := "Fire"
    result := typeCase.Type.Name

    if result != expected {
        t.Errorf("Expected Type Name %s; got %s", expected, result)
    }
}
