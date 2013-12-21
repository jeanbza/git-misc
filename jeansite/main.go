package main

import (
    //"html/template"
    "io/ioutil"
    "net/http"
    //"regexp"
    "os"
    "fmt"
)

func main() {
    http.HandleFunc("/", indexPage)

    fileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    checkError(err)
}

func indexPage(rw http.ResponseWriter, req *http.Request) {
    index, _ := ioutil.ReadFile("html/index.html")
    rw.Write([]byte(index))
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}