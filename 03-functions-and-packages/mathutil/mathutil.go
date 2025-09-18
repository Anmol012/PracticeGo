package mathutil

// Add is an exported function (public) FunctionName first char is Upper case 
func Add(a int, b int) int {
    return a + b
}

// subtract is an unexported function (private) functionName first char is lower case 
func subtract(a int, b int) int {
    return a - b
}
