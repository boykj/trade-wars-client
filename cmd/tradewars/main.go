package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	gotdotenv.load()
	Port := os.Getenv('PORT')
	log.Println("Starting server on :"+PORT)
	err := http.ListenAndServe(";"+PORT, mux)
	log.Fatal(err)
}