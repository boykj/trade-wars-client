package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	mux := http.NewServeMux()
	godotenv.Load()
	PORT := os.Getenv("PORT")
	log.Println("Starting server on :"+PORT)

	mux.Handle("/", http.HandlerFunc(home))
	
	http.HandleFunc("/", home)
	http.HandleFunc("/game/", game)
	http.HandleFunc("/chat/", chat)
	http.HandleFunc("/trade/", trade)

	err := http.ListenAndServe(":"+PORT, nil)
	log.Fatal(err)
}