package main

import (
	"encoding/json"
	"log"
	"math"
	"mime"
	"net/http"
	"time"

	bfg "github.com/kimle/bfg/pkg/generator"
	"github.com/rs/cors"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}

func getIngredients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	meal := bfg.Generate()
	if err := json.NewEncoder(w).Encode(meal); err != nil {
		log.Printf("could not encode meal: %v", err)
	}
}

func getMeekTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		loc = time.UTC
		log.Printf("could not load tzdata: %v\n", err)
	}
	courtdate := time.Date(2017, time.November, 6, 0, 0, 0, 0, loc)
	daysSince := math.Floor(time.Since(courtdate).Hours() / (time.Hour * 24).Hours())
	err = json.NewEncoder(w).Encode(map[string]interface{}{"days": daysSince})
	if err != nil {
		log.Fatalf("app must have message about freeing meek mill: %v\n", err)
	}
}

func main() {
	bfg.SetUp()
	if err := mime.AddExtensionType(".js", "application/javascript; charset=utf-8"); err != nil {
		log.Println(err)
	}
	if err := mime.AddExtensionType(".css", "text/css; charset=utf-8"); err != nil {
		log.Println(err)
	}
	log.Println("listening on port 8080...")
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("dist/static"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/add", getIngredients)
	mux.HandleFunc("/freemeek", getMeekTime)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
