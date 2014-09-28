package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "regexp"
)

/*
plan:
STAGE 1
1. save page locally
2. save page javascript+css locally
3. change page javascript+css to local paths

STAGE 2
1. spawn workers to repeat STAGE1 for each link on page
2. change links to local paths

STAGE 3
1. allow for ability to recurse n times
2. add config files ot describe recursive length
*/

func main() {
    response, err := http.Get("http://localhost:8080")
    handleErr(err)

    bodyByte, err := ioutil.ReadAll(response.Body)
    handleErr(err)

    bodyStr := string(bodyByte)

    fmt.Println("Before")
    outputLinks(bodyStr)

    bodyStrReplaced := replaceLinks(bodyStr)

    fmt.Println("After")
    outputLinks(bodyStrReplaced)
}

func replaceLinks(bodyStr string) string {
    hrefRegex := regexp.MustCompile(`(href|src)="(https://www)?(.*)"`)
    result := hrefRegex.ReplaceAllString(bodyStr, "$1=\"localhost:1234$3\"")

    return result
}

func outputLinks(bodyStr string) {
    hrefRegex := regexp.MustCompile(`(href|src)=".*"`)
    result := hrefRegex.FindAllString(bodyStr, -1)

    for _, val := range result {
        fmt.Println(val)
    }
}

func handleErr(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
}