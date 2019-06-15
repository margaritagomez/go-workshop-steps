package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	pokemons []Pokemon
	pokedex  map[int]Pokemon
)

func init() {
	newPokemon(1, "Bulbasaur", "Ivysaur", []string{"grass", "venom"})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
}
