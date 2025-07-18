package main
import "fmt"

func main ()	{
	// Relational and Logical Operators

	// Relational Operators are ==, !=, <, >, <=, >=
	var age int = 18

	var isAdult bool = age > 18
	fmt.Println("Is adult:", isAdult) 

	// Logical Operators are &&, ||, !
	var isStudent bool = true
	var isEmployed bool = false
	var isEligibleForLoan bool = isAdult && !isStudent && isEmployed
	fmt.Println("Is eligible for loan:", isEligibleForLoan)

	// Statements vs Expressions in Go
	// A statement is a complete instruction that performs an action.
	var x int = 10
	var y int = 20
	var sum int = x + y // This is an expression
	fmt.Println("Sum:", sum)
	var isEqual bool = (x == y) // This is also an expression
	fmt.Println("Is x equal to y:", isEqual)

	// An expression is a combination of variables, constants, operators, and function calls that evaluates to a value.
	var isGreater bool = x > y // This is an expression
	fmt.Println("Is x greater than y:", isGreater)
	// A statement can contain expressions, but an expression cannot stand alone as a statement.
	var message string = "Hello, Go!" // This is a statement
	fmt.Println(message) // This is a statement that uses the expression message
	// An expression can be part of a statement, but it cannot be a standalone statement.
	var result int = (x + y) * 2 // This is a statement that contains an expression
	fmt.Println("Result:", result)
}