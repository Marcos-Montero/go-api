package pokemon

// Pokemon represents the basic details of a Pokémon
type Pokemon struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Types []PokemonType `json:"types"`
    // Add other relevant fields as needed
}

// PokemonType represents the type of a Pokémon, such as "Grass", "Fire", etc.
type PokemonType struct {
    Slot int    `json:"slot"`
    Type struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"type"`
}

// NewPokemon creates and returns a new instance of a Pokemon with the given details
func NewPokemon(id int, name string, types []PokemonType) *Pokemon {
    return &Pokemon{
        ID:    id,
        Name:  name,
        Types: types,
    }
}
