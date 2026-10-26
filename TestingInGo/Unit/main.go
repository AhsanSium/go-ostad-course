package main

import "fmt"

func isPrime(n int) (bool, string) {
	if n <= 1 {
		return false, "not a prime number"
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, "not a prime number"
		}
	}
	return true, "is a prime number"
}

func main() {
	n := 7
	isPrime1, msg1 := isPrime(n)
	fmt.Println(n, msg1)
	if !isPrime1 {
		fmt.Println(n, "is not a prime number")
	} else {
		fmt.Println(n, "is a prime number")
	}
	n2 := 10
	isPrime2, msg2 := isPrime(n2)
	fmt.Println(n2, msg2)
	if !isPrime2 {
		fmt.Println(n2, "is not a prime number")
	} else {
		fmt.Println(n2, "is a prime number")
	}
}
