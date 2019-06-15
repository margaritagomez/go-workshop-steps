package main

import (
	"fmt"
	"net/http"
)

func getPokemons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hola")
}
