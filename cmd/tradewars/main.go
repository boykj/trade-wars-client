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
<<<<<<< HEAD
	http.HandleFunc("/", index)
=======

	mux.Handle("/", http.HandlerFunc(home))
	
	http.HandleFunc("/", home)
>>>>>>> fec322344a9f6d5518da6a46c3b1d39695e6208c
	http.HandleFunc("/game/", game)
	http.HandleFunc("/chat/", chat)
	http.HandleFunc("/trade/", trade)

	err := http.ListenAndServe(":"+PORT, nil)
	log.Fatal(err)
}