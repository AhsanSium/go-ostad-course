package main

import "testing"

func Test_isPrime(t *testing.T) {
	isPrime1, msg1 := isPrime(0)
	if !isPrime1 {
		t.Errorf("Expected 7 to be a prime number, but got %s", msg1)
	}
	isPrime2, msg2 := isPrime(10)
	if isPrime2 {
		t.Errorf("Expected 10 to not be a prime number, but got %s", msg2)
	}

}