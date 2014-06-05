package main

import "fmt"

func main() {
    var input = ""

    for input != "q" && input != "quit" {
        _, err := fmt.Scanf("%s", &input)
        if (err != nil) {
            
        }

        fmt.Printf("Read %v from stdin\n", input)
        fmt.Printf("%v\n", OrderCoffee(5))
    }
}

// Orders amountOfCoffees coffees and returns the price
// when the coffees are ready
func OrderCoffee(amountOfCoffees int) (price int) {
    return amountOfCoffees*3
}