package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var pokemons []Pokemon

func init() {
	pokemons = []Pokemon{
		Pokemon{
			PokedexNumber: 1,
			Name:          "Bulbasaur",
			EvolvesTo:     "Ivysaur",
			Types: []Type{
				TypeGrass,
				TypePoison,
			},
		},
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
