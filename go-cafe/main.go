package main

import "fmt"

func main() {
    var input = ""

    for input != "q" && input != "quit" {
        _, err := fmt.Scanf("%s", &input)
        if (err != nil) {
            
        }

        fmt.Printf("Read %v from stdin\n", input)
    }
}