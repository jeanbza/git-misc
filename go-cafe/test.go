package main

import "fmt"

func main() {
	foo := "bar"
	go func test() {
		fmt.Println(foo)
	}()
}

