package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	mux := http.NewServeMux()
	// mux.HandleFunc("/", redirect)	
	mux.Handle("/", http.HandlerFunc(playersHandler))
	mux.Handle("/map", http.HandlerFunc(mapHandler))

	fileServer := http.FileServer(http.Dir(".ui/static"))
	mux.Handle("/static", http.StripPrefix("/static", fileServer))

	godotenv.Load()

	PORT := os.Getenv("PORT")
	log.Println("Starting server on :"+PORT)
	//http.ListenAndServe(":"+PORT, mux)
	err := http.ListenAndServe(":"+PORT, nil)
	log.Fatal(err)
}
