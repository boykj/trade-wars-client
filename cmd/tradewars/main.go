package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	mux := http.NewServeMux()
	mux.HandlerFunc("/", home)
	mux.HandlerFunc("/players", playerHandler)
	mux.HandlerFunc("/map", mapHandler)

	fileServer := http.FileServer(http.Dir(".ui/static"))
	mux.Handle("/static", http.StripPrefix("/static", fileServer))

	godotenv.Load()

	PORT := os.Getenv("PORT")
	log.Println("Starting server on :"+PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
