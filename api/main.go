package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	pokemons []*Pokemon
	pokedex  = make(map[int]*Pokemon)
)

func init() {
	newPokemon(1, "Bulbasaur", "Ivysaur", 25, []string{"grass", "venom"})
	newPokemon(7, "Squirtle", "Wartortle", 14, []string{"water"})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemons", getMyPokemons).Methods("GET")
	router.HandleFunc("/pokemons", capturePokemon).Methods("POST")
	router.HandleFunc("/pokemons/{pokedex_number}", choosePokemon).Methods("GET")
	router.HandleFunc("/pokemons/{pokedex_number}", trainPokemon).Methods("PUT")
	router.HandleFunc("/pokemons/{pokedex_number}", transferPokemon).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
