package main

import (
	"encoding/json"
	"log"
	"mime"
	"net/http"

	"github.com/kimle/bfg/bfg"
	"github.com/rs/cors"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func addIngredient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	meal := bfg.Generate()
	err := json.NewEncoder(w).Encode(meal)
	if err != nil {
		log.Fatal(err)
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
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/add", addIngredient)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
