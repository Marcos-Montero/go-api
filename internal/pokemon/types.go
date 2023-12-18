package pokemon

// Pokemon represents the basic details of a Pokémon
type PokemonType struct {
	Slot int    `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type PokemonAbility struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
	Ability  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"ability"`
}

type PokemonStat struct {
	Stat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
	Effort int `json:"effort"`
	Base   int `json:"base"`
}

type PokemonMove struct {
	Move struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move"`
	VersionGroupDetails []struct {
		LevelLearnedAt  int `json:"level_learned_at"`
		VersionGroup    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
	} `json:"version_group_details"`
}

type PokemonSprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type Pokemon struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Types      []PokemonType    `json:"types"`
	Abilities  []PokemonAbility `json:"abilities"`
	Stats      []PokemonStat    `json:"stats"`
	Moves      []PokemonMove    `json:"moves"`
	Sprites    PokemonSprites   `json:"sprites"`
	Height     int              `json:"height"`
	Weight     int              `json:"weight"`
	BaseExperience int          `json:"base_experience"`
	Order      int              `json:"order"`
	IsDefault  bool             `json:"is_default"`
}

func NewPokemon(id int, name string, types []PokemonType) *Pokemon {
	return &Pokemon{
		ID:    id,
		Name:  name,
		Types: types,
	}
}
