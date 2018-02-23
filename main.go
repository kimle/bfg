package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "svamp och majs")
}

func main() {
	fmt.Printf("listening on port %s...", port)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(port, nil))
}
