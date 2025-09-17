package main

import "fmt"

func main() {
    // Multiple variable declaration
    var x, y, z int = 10, 20, 30

    // Mixed multiple types
    var (
        name   = "GoLang"
        version float32 = 1.21
        stable  = true
    )

    fmt.Println("x:", x, "y:", y, "z:", z)
    fmt.Println("Name:", name, "Version:", version, "Stable:", stable)
}
