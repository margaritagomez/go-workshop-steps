package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getMyPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func choosePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get pokemon")
}

func capturePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create pokemon")
}

func renamePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update pokemon")
}

func transferPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete pokemon")
}
