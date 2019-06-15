package main

type Pokemon struct {
	PokedexNumber int    `json:"pokedex_number"`
	Name          string `json:"name"`
	EvolvesTo     string `json:"evolves_to,omitempty"`
	Types         []Type `json:"types"`
}

type Type string

const (
	TypeGrass    Type = "Grass"
	TypeDark     Type = "Dark"
	TypePsychic  Type = "Psychic"
	TypeFighting Type = "Fighting"
	TypeIce      Type = "Ice"
	TypeGround   Type = "Ground"
	TypeElectric Type = "Electric"
	TypeBug      Type = "Bug"
	TypeWater    Type = "Water"
	TypeRock     Type = "Rock"
	TypeFairy    Type = "Fairy"
	TypeFlying   Type = "Flying"
	TypeFire     Type = "Fire"
	TypePoison   Type = "Poison"
	TypeNormal   Type = "Normal"
	TypeGhost    Type = "Ghost"
	TypeDragon   Type = "Dragon"
	TypeSteel    Type = "Steel"
)
