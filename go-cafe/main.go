package main

import (
    "fmt"
    "os"
    "time"
    "math/rand"
)

func main() {
    costs := map[string]int {
        "Coffee": 5,
        "Tea": 4,
        "Hot Chocolate": 1,
    }

    orders := getOrders()
    fmt.Println("I'll let you know as soon as those orders are ready!\n")
    
    // Spawn goroutine for each order
    // (todo) spawn goroutines capped at max amount of barristas
    ch, totalOrders := completeOrders(orders)

    for i := 0; i < totalOrders; i++ {
        fmt.Println(<- ch)
    }

    // Tally total cost and give customer their order
    cost := tallyCost(orders, costs)
    fmt.Printf("\nHere's your order! Your total comes to $%d. Have a good one!\n", cost)
}

func completeOrders(orders map[string]int) (<- chan string, int) {
    ch := make(chan string)
    totalOrders := 0

    for product, amt := range orders {
        for i := 0; i < amt; i++ {
            totalOrders++
            go func(product string) {
                amt := time.Duration(rand.Intn(1000))
                time.Sleep(time.Millisecond * amt)
                ch <- fmt.Sprintf("%v completed!", product)
            }(product)
        }
    }

    return ch, totalOrders
}

func tallyCost(orders map[string]int, costs map[string]int) (int) {
    cost := 0

    for product, amt := range orders {
        cost += amt * costs[product]
    }

    return cost
}

func getOrders() (map[string]int) {
    var input = ""

    orders := map[string]int{
        "Coffee": 0,
        "Tea": 0,
        "Hot Chocolate": 0,
    }

    fmt.Println("Welcome! What can I get you? We have coffee (c), hot chocolate (h), and tea (t). If you'd like to quit, press (q).")

    for input != "d" && input != "done" {
        _, err := fmt.Scanf("%s", &input)
        if (err != nil) {
            
        }

        switch input {
            case "c", "coffee":
                fmt.Println("One coffee, got it.")
                orders["Coffee"]++
            case "h", "hot chocolate":
                fmt.Println("One hot chocolate, got it.")
                orders["Hot Chocolate"]++
            case "t", "tea":
                fmt.Println("One tea, got it.")
                orders["Tea"]++
            case "q", "quit":
                os.Exit(1)
            case "d", "done":
            default:
                fmt.Println("Huh?")
        }

        if input != "d" && input != "done" {
            fmt.Println("What else can I get you?  We have coffee (c), hot chocolate (h), and tea (t). Or, are you done (d)?")
        }
    }

    fmt.Println("That's it? Ok!")

    return orders
}