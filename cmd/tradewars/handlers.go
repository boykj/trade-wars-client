package main

import (
    "fmt"        
    "net/http"
    "strconv"
    "log"
    "html/template"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        w.WriteHeader(405)
        w.Write([]byte("Method not allowed"))
        return
    }

    w.Write([]byte("Create a new snippet..."))
}

func game(w http.ResponseWriter, r *http.Request) {
    template.ParseFiles("./ui/html/game.page.tmpl")
}

func chat(w http.ResponseWriter, r *http.Request) {
    template.ParseFiles("./ui/html/chat.page.tmpl")
}

func trade(w http.ResponseWriter, r *http.Request) {
    template.ParseFiles("./ui/html/trade.page.tmpl")
}



