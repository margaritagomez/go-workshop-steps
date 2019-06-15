package main

import (
	"fmt"
	"net/http"
)

func getMyPokemons(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get pokemons")
}

func choosePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get pokemon")
}

func capturePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create pokemon")
}

func evolvePokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update pokemon")
}

func transferPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete pokemon")
}
