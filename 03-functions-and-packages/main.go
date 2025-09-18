package main

import (
    "fmt"
    "github.com/Anmol012/golang-practice/03-functions-and-packages/mathutil"
)

func main() {
    sum := mathutil.Add(10, 20)
    fmt.Println("Sum:", sum)
}