package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Calculator for basic arithmetic operations
	// var num1 float64
	// var num2 float64

	// Using fmt.Scanf to read user input
	// fmt.Print("Enter first number: ")
	// fmt.Scanf("%f", &num1)
	// fmt.Print("Enter second number: ")
	// fmt.Scanf("%f", &num2)

	// Using fmt.Scanln to read user input
	// fmt.Print("Enter first number and second number: ")
	// fmt.Scanln(&num1, &num2)
	// var sum float64 = num1 + num2

	// fmt.Println("Sum:", sum)


	// Full Name input	var firstName string
	// var firstName string
	// var lastName string
	// fmt.Print("Enter your first name and last name: ")
	// // fmt.Scanf("%s %s", &firstName, &lastName)
	// fmt.Scanln(&firstName, &lastName)

	// var fullName string = firstName + " " + lastName
	// fmt.Printf("Your full name is: %s\n", fullName)

	// Taking input using bufio
	// This is useful for reading strings with spaces
	var city string
	fmt.Print("Enter your city: ")
	reader := bufio.NewReader(os.Stdin)
	cityInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	city = cityInput[:len(cityInput)-1] // Remove the newline character
	fmt.Printf("You live in: %s\n", city)
}