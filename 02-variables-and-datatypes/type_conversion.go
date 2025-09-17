package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. Integer to Float and vice versa
	var i int = 42
	var f float64 = float64(i) // convert int -> float64
	var j int = int(f)         // convert float64 -> int decimal truncated

	fmt.Println("Int to Float:", f)
	fmt.Println("Float back to Int:", j)

	// 2. Integer to String
	var num int = 100
	var str string = strconv.Itoa(num) // using strconv.Itoa
	fmt.Println("Int to String:", str)

	// 3. String to Integer
	strNum := "250"
	numConv, err := strconv.Atoi(strNum) // using strconv.Atoi
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("String to Int:", numConv)
	}

	// 4. Float to String
	var pi float64 = 3.14159
	piStr := fmt.Sprintf("%.2f", pi) // formatted float to string
	fmt.Println("Float to String:", piStr)

	// 5. String to Float
	strFloat := "12.34"
	floatConv, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println("String to Float:", floatConv)
	}

	// 6. Boolean to String
	var isActive bool = true
	boolStr := strconv.FormatBool(isActive)
	fmt.Println("Bool to String:", boolStr)

	// 7. String to Boolean
	strBool := "false"
	boolConv, err := strconv.ParseBool(strBool)
	if err != nil {
		fmt.Println("Error converting string to bool:", err)
	} else {
		fmt.Println("String to Bool:", boolConv)
	}

	// 8. Rune (char) to String
	var r rune = 'A'
	runeToStr := string(r)
	fmt.Println("Rune to String:", runeToStr)

	// 9. String to Byte Slice (and back)
	text := "Hello"
	byteSlice := []byte(text)
	strFromBytes := string(byteSlice)
	fmt.Println("String to Bytes:", byteSlice)
	fmt.Println("Bytes to String:", strFromBytes)
}
