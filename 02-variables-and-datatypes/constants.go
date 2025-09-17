package main

import "fmt"

func main() {

	// you can not change constant values
    const PI = 3.14159
    const Greeting string = "Hello!"

    fmt.Println("PI:", PI)
    fmt.Println("Greeting:", Greeting)

    // Multiple constants
    const (
        a = 10
        b = 20
    )
    fmt.Println("a:", a, "b:", b)
}
