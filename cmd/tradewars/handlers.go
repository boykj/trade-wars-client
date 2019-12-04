package main

import (    
    "net/http"
    "log"
    "html/template"
    "time"
)

func home(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        "./ui/html/home.page.tmpl",
        "./ui/html/base.layout.tmpl",
    }

    ts, err := template.ParseFiles(files...)

    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.Execute(w, "ui/html/home.page.tmpl")

    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
    
    if r.Method==http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
            return   
        }
        callsign := r.Form.Get("callsign")
        cookie := http.Cookie {
            Name: "callsign",
            Value: callsign,
            Expires: time.Now().AddDate(0, 0, 1),
            Path: "/",
        }
        http.SetCookie(w, &cookie)
        http.Redirect(w, r, "/map", http.StatusSeeOther)

    } else {
        http.Error(w, "Method Not Allowed", 405)
        return
    }
}

 func mapHandler(w http.ResponseWriter, r *http.Request) {
    files := []string{
    "./ui/html/game.page.tmpl",
    "./ui/html/base.layout.tmpl",
    }

    ts, err := template.ParseFiles(files...)
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
    var cookie, errC = r.Cookie("callsign")
    if errC != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error || Cookie callsign not resolved", 500)
        return
    }
    callsign := cookie.Value
    
    w.Write([]byte("Welcome to Trade wars, " + callsign))
    w.Write([]byte(" Create a new map..."))
 }

 func chatHandler(w http.ResponseWriter, r *http.Request) {
    files := []string{
        "./ui/html/chat.page.tmpl",
        "./ui/html/base.layout.tmpl",
    }

    ts, err := template.ParseFiles(files...)
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
    var cookie, errC = r.Cookie("callsign")
    if errC != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error || Cookie callsign not resolved", 500)
        return
    }
    callsign := cookie.Value
        w.Write([]byte("Welcome to Trade wars Chat, " + callsign))
 }

 func tradeHandler(w http.ResponseWriter, r *http.Request) {
    files := []string{
        "./ui/html/trade.page.tmpl",
        "./ui/html/base.layout.tmpl",
    }

    ts, err := template.ParseFiles(files...)
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
    var cookie, errC = r.Cookie("callsign")
    if errC != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error || Cookie callsign not resolved", 500)
        return
    }
    callsign := cookie.Value
        w.Write([]byte("Welcome to Trade wars Trade Hub, " + callsign))
 }

 func redirect(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        http.Redirect(w, r, "/players", 303)
    }
}
