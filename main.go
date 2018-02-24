package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// The core ingredients for an exit fragger meal
const (
	Svamp     = "svamp"
	Majs      = "majs"
	Paprika   = "paprika"
	Makaroner = "makaroner"
	Hona      = "höna"
	Kott      = "kött"
	Vitsas    = "vit sås"
)

var ingredients []string
var meal map[string]struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
	ingredients = append(ingredients, Svamp, Majs, Paprika, Makaroner, Hona, Kott, Vitsas)
	meal = make(map[string]struct{})
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "the complete dish: ")
	if err != nil {
		log.Fatal(err)
	}
	if len(meal) > 0 {
		for key := range meal {
			_, err := fmt.Fprintf(w, "%v ", key)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func addIngredient(w http.ResponseWriter, r *http.Request) {
	if len(meal) >= 3 {
		meal = make(map[string]struct{})
		return
	}
	for {
		ingredient := ingredients[rand.Intn(len(ingredients))]
		if _, ok := meal[ingredient]; !ok {
			meal[ingredient] = struct{}{}
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	log.Println("listening on port 8080...")
	http.HandleFunc("/", index)
	http.HandleFunc("/add", addIngredient)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
