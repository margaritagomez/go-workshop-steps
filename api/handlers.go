package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getMyPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func choosePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	pokedexNumber, err := strconv.Atoi(params["pokedex_number"])
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, p := range pokemons {
		if p.PokedexNumber == pokedexNumber {
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(p)
			if err != nil {
				fmt.Println(err.Error())
			}

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Pokemon{})
}

func capturePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newPokemon Pokemon

	err := json.NewDecoder(r.Body).Decode(&newPokemon)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, ok := pokedex[newPokemon.PokedexNumber]
	if ok {
		http.Error(w, "Pokemon already exists", 400)
		return
	}

	pokemons = append(pokemons, &newPokemon)
	pokedex[newPokemon.PokedexNumber] = &newPokemon

	json.NewEncoder(w).Encode(&newPokemon)
}

func trainPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	pokedexNumber, err := strconv.Atoi(params["pokedex_number"])
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, p := range pokemons {
		if p.PokedexNumber == pokedexNumber {
			p.Level = p.Level + 1

			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(p)
			if err != nil {
				fmt.Println(err.Error())
			}

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Pokemon{})
}

func transferPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	pokedexNumber, err := strconv.Atoi(params["pokedex_number"])
	if err != nil {
		fmt.Println(err.Error())
	}

	for i, p := range pokemons {
		if p.PokedexNumber == pokedexNumber {
			pokemons = append(pokemons[:i], pokemons[i+1:]...)
			delete(pokedex, pokedexNumber)

			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(p)
			if err != nil {
				fmt.Println(err.Error())
			}

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Pokemon{})
}
