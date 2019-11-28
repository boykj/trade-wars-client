package main

import (    
    "net/http"
    "log"
    "html/template"
    // "time"
)

func playersHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    // GET Method 
    if r.Method==http.MethodGet{
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

        err = ts.Execute(w, nil)
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
        }
        
    // POST Method
    } else if r.Method==http.MethodPost {
        if err := r.ParseForm(); err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
            return
        }
        callsign := r.Form.Get("callsign")
        cookie := http.Cookie{
            Name: "callsign",
            Value: callsign,
            // MaxAge: 60 * 60, // For usage without time import
            Expires: time.Now().AddDate(0, 0, 1),
            Path: "/",

        }
        http.SetCookie(w, &cookie)

    } else {
        http.Error(w, "Method Not Allowed", 405)
        return
    }
}

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/map", 303)
}

 func mapHandler(w http.ResponseWriter, r *http.Request) {
    var cookie, err = r.Cookie("callsign")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }
    callsign := cookie.Value
    log.Println("Welcome to Trade wars, " + callsign)

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
    w.Write([]byte("Create a new map..."))
 }

//  func chat(w http.ResponseWriter, r *http.Request) {
//     if r.URL.Path != "/" {
//         http.NotFound(w, r)
//         return
//     }

//     files := []string{
//         "./ui/html/chat.page.tmpl",
//         "./ui/html/base.layout.tmpl",
//     }

//     ts, err := template.ParseFiles(files...)
//     if err != nil {
//         log.Println(err.Error())
//         http.Error(w, "Internal Server Error", 500)
//         return
//     }

//     err = ts.Execute(w, nil)
//     if err != nil {
//         log.Println(err.Error())
//         http.Error(w, "Internal Server Error", 500)
//     }
//  }

//  func trade(w http.ResponseWriter, r *http.Request) {
//     if r.URL.Path != "/" {
//         http.NotFound(w, r)
//         return
//     }

//     files := []string{
//         "./ui/html/trade.page.tmpl",
//         "./ui/html/base.layout.tmpl",
//     }

//     ts, err := template.ParseFiles(files...)
//     if err != nil {
//         log.Println(err.Error())
//         http.Error(w, "Internal Server Error", 500)
//         return
//     }

//     err = ts.Execute(w, nil)
//     if err != nil {
//         log.Println(err.Error())
//         http.Error(w, "Internal Server Error", 500)
//     }
//  }

//  func redirect(w http.ResponseWriter, r *http.Request) {
//     if r.URL.Path == "/" {
//         http.Redirect(w, r, "/players", 303)
//     }
// }
