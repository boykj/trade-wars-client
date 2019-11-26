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

func game(w http.ResponseWriter, r *http.Request) {
    template.ParseFiles("./ui/html/game.page.tmpl")
}

func chat(w http.ResponseWriter, r *http.Request) {
    template.ParseFiles("./ui/html/chat.page.tmpl")
}

func trade(w http.ResponseWriter, r *http.Request) {
    template.ParseFiles("./ui/html/trade.page.tmpl")
}



