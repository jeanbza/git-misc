package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
}

func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/save", savePage)

    fileServer := http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }
}

func homePage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "My Home",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["home.html"] = template.Must(template.ParseFiles("html/home.html", "html/layout.html"))
    tmpl["home.html"].ExecuteTemplate(rw, "base", p)
}

func savePage(rw http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    firstName := req.PostFormValue("first_name")
    lastName := req.PostFormValue("last_name")
    fmt.Println(firstName)
    fmt.Println(lastName)

    http.Redirect(rw, req, "/", http.StatusFound)
}