package main

type Type string

const (
	TypeGrass    Type = "grass"
	TypeDark     Type = "dark"
	TypePsychic  Type = "psychic"
	TypeFighting Type = "fighting"
	TypeIce      Type = "ice"
	TypeGround   Type = "ground"
	TypeElectric Type = "electric"
	TypeBug      Type = "bug"
	TypeWater    Type = "water"
	TypeRock     Type = "rock"
	TypeFairy    Type = "fairy"
	TypeFlying   Type = "flying"
	TypeFire     Type = "fire"
	TypePoison   Type = "poison"
	TypeNormal   Type = "normal"
	TypeGhost    Type = "ghost"
	TypeDragon   Type = "dragon"
	TypeSteel    Type = "steel"
)

type Pokemon struct {
	PokedexNumber int    `json:"pokedex_number"`
	Name          string `json:"name"`
	EvolvesTo     string `json:"evolves_to,omitempty"`
	Types         []Type `json:"types"`
	Level         int    `json:"level"`
}

func newPokemon(pokedexNumber int, name string, evolvesTo string, level int, types []string) *Pokemon {
	newPokemon := &Pokemon{
		PokedexNumber: pokedexNumber,
		Name:          name,
		EvolvesTo:     evolvesTo,
		Level:         level,
	}

	for _, t := range types {
		newPokemon.Types = append(newPokemon.Types, Type(t))
	}

	pokemons = append(pokemons, newPokemon)
	pokedex[pokedexNumber] = newPokemon

	return newPokemon
}
