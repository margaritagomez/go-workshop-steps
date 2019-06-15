package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	pokemons []Pokemon
	pokedex  = make(map[int]Pokemon)
)

func init() {
	newPokemon(1, "Bulbasaur", "Ivysaur", []string{"grass", "venom"})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemons", getMyPokemons).Methods("GET")
	router.HandleFunc("/pokemons", capturePokemon).Methods("POST")
	router.HandleFunc("/pokemons/{pokedexID}", choosePokemon).Methods("GET")
	router.HandleFunc("/pokemons/{pokedexID}", renamePokemon).Methods("PUT")
	router.HandleFunc("/pokemons/{pokedexID}", transferPokemon).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
