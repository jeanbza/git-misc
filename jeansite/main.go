package main

import (
    "html/template"
    // "io/ioutil"
    "net/http"
    "os"
    "fmt"
)

type Post struct {
    Title   string
    Content string
}

func main() {
    http.HandleFunc("/", indexPage)

    fileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    checkError(err)
}

func indexPage(rw http.ResponseWriter, req *http.Request) {
    p := []Post{
        Post{
            Title: "Test blog title 1", 
            Content: "Test blog content 1",
        },
        Post{
            Title: "Test blog title 2",
            Content: "Test blog content 2",
        },
        Post{
            Title: "Test blog title THE LAST",
            Content: "Test blog content BOOM THE LAST",
        },
    }

    t, _ := template.ParseFiles("html/index.html")
    t.Execute(rw, p)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}