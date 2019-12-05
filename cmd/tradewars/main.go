package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/players", playerHandler)
	mux.HandleFunc("/chat", chatHandler)
	mux.HandleFunc("/trade", tradeHandler)
	mux.HandleFunc("/map", mapHandler)
	mux.HandleFunc("/ws", handleConnections)
	//go handleMessages()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	godotenv.Load()

	PORT := os.Getenv("PORT")
	log.Println("Starting server on :"+PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
