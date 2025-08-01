package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Type Conversion in GO
	var age int = 25
	var floatAge float64 = float64(age) // int to float64
	fmt.Printf("Age: %d, Type: 	%T\n", age, age)
	fmt.Printf("Float Age: %f, Type: 	%T\n", floatAge, floatAge)

	var pi float64 = 3.1416
	var intPi int = int(pi) // float64 to int
	fmt.Printf("Pi: %f, Type: 	%T\n", pi, pi)
	fmt.Printf("Int Pi: %d, Type: 	%T\n", intPi, intPi)

	var stringValue string = "123"
	var intValue, err = strconv.Atoi(stringValue) // string to int
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
	} else {
		fmt.Printf("String Value: %s, Type: 	%T\n", stringValue, stringValue)
		fmt.Printf("Int Value: %d, Type: 	%T\n", intValue, intValue)
	}

	var intValue2 int = 456
	var stringValue2 string = strconv.Itoa(intValue2) // int to string
	fmt.Printf("int value 2 : %d, Type: %T\n", intValue2, intValue2)
	fmt.Printf("String Value 2: %s, Type: %T\n", stringValue2, stringValue2)
}
