package bfg

import (
	"math/rand"
	"sort"
	"time"
)

// The ingredients for a true exit fragger experience
const (
	Svamp     = "Svamp"
	Majs      = "Majs"
	Paprika   = "Paprika"
	Makaroner = "Makaroner"
	Hona      = "Höna"
	Kott      = "Kött"
	Vitsas    = "Vit sås"
	TortillaP = "Tortilla plättar"
)

var ingredients []string
var meal map[string]struct{}

// Meal is the struct to be JSON encoded and sent to the client
type Meal struct {
	Ingredients []string `json:"ingredients"`
}

// SetUp sets up by populating the ingredients slice etc.
func SetUp() {
	rand.Seed(time.Now().UnixNano())
	ingredients = append(ingredients, Svamp, Majs, Paprika, Makaroner, Hona, Kott, Vitsas, TortillaP)
	meal = make(map[string]struct{})
}

// Generate returns a struct of type Meal, which contains a sorted slice.
func Generate() Meal {
	var dish Meal
	if len(meal) >= 4 {
		meal = make(map[string]struct{})
	}
	for {
		if len(meal) >= 4 {
			for ingr := range meal {
				dish.Ingredients = append(dish.Ingredients, ingr)
			}
			break
		}

		ingredient := ingredients[rand.Intn(len(ingredients))]
		if _, ok := meal[ingredient]; !ok {
			meal[ingredient] = struct{}{}
		}
	}
	sort.Strings(dish.Ingredients)
	return dish
}
