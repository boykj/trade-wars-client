package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	// mux := http.NewServeMux()
	godotenv.Load()
	PORT := os.Getenv("PORT")
	log.Println("Starting server on :"+PORT)
	http.HandleFunc("/", index)
	http.HandleFunc("/game/", game)
	http.HandleFunc("/chat/", chat)
	http.HandleFunc("/trade/", trade)
	http.ListenAndServe(":"+PORT, nil)
	// err := http.ListenAndServe(":"+PORT, nil)
	// log.Fatal(err)
}