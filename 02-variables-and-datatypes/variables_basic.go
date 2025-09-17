package main

import "fmt"

func main() {
    // declaring with var
    var name string = "Anmol"
    var age int = 25
    var isEngineer bool = true

    // Go infers the type automatically
    var city = "Delhi"

    // Short declaration (only inside functions)
    salary := 50000.50

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Engineer:", isEngineer)
    fmt.Println("City:", city)
    fmt.Println("Salary:", salary)
}
