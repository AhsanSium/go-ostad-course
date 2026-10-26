package main

import "testing"

func Test_isPrime(t *testing.T) {
	// isPrime1, msg1 := isPrime(0)
	// if !isPrime1 {
	// 	t.Errorf("Expected 7 to be a prime number, but got %s", msg1)
	// }
	// isPrime2, msg2 := isPrime(10)
	// if isPrime2 {
	// 	t.Errorf("Expected 10 to not be a prime number, but got %s", msg2)
	// }

	tests := []struct {
		n int
		expected string
	}{
		// {1, "is a prime number"},
		{7, "is a prime number"},
		{10, "not a prime number"},
		{0, "not a prime number"},
		{1, "not a prime number"},
		{2, "is a prime number"},
		{3, "is a prime number"},
		{4, "not a prime number"},
	}

	for _, test := range tests {
		t.Run("Testing isPrime function", func(t *testing.T) {
			_, msg := isPrime(test.n)
			if msg != test.expected {
				t.Errorf("Expected %d to be %s, but got %s", test.n, test.expected, msg)
			}
		})
	}

}