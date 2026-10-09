package main

import "fmt"

/*
1. Any type that has all the methods listed in an interface definition is said to implement that interface. In Go, you don't need to explicitly declare that a type implements an interface; it happens automatically if the type has the required methods.

2. A type that satisfy an interface can be used anywhere that interface is expected. This allows for polymorphism, where different types can be treated uniformly through the interface they implement.
*/

type MyInterface interface {
	MethodWithoutParameters ()
	MethodWithParameters (int)
	MethodWithReturnValue () string
}

type MyType string

func main () {
	myType := MyType("Hello, World!")
	
	var myInterface MyInterface

	myInterface = myType // Assigning myType to myInterface

	myInterface.MethodWithoutParameters()
	myInterface.MethodWithParameters(42)
	result := myInterface.MethodWithReturnValue()

	fmt.Println("my interface:", myInterface, "my type:", myType)
	fmt.Println("result:", result)
}

func (m MyType) MethodWithoutParameters () {}
func (m MyType) MethodWithParameters (param int) {}
func (m MyType) MethodWithReturnValue () string {
	return string(m)
}