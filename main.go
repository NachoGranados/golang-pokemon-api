package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokemon-api/database"
	"testing"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDbAsValueArray())
}

func addPokemon(w http.ResponseWriter, r *http.Request) {
	var newPokemon database.Pokemon

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newPokemon)

	if _, ok := database.PokemonDb[newPokemon.ID]; ok {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	database.PokemonDb[newPokemon.ID] = newPokemon
	w.WriteHeader(http.StatusOK)
}

func handleRequests() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons", addPokemon).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}




func Sum(a, b int) int {
	return a + b
}



// Test methods start with Test
func TestSum(t *testing.T) {
	// Note that the data variable is of type array of anonymous struct,
	// which is very handy for writing table-driven unit tests.
	data := []struct {
		a, b, res int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{1, -1, 0},
		{2, 3, 5},
		{1000, 234, 1234},
	}

	for _, d := range data {
		if got := Sum(d.a, d.b); got != d.res {
			t.Errorf("Sum(%d, %d) == %d, want %d", d.a, d.b, got, d.res)
		}
	}
}



























func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}

