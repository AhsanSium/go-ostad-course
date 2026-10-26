/*
We have a counter and two go routines that are incrementing the counter. The problem is that both goroutines are trying to read and write to the counter at the same time, which can lead to a race condition. A race condition occurs when two or more goroutines access shared data and try to change it at the same time. This can lead to unexpected behavior and bugs in your program.
*/
package main

import (
	"fmt"
	"sync"
)
var counter = 0
var mu sync.Mutex // Mutex is a mutual exclusion lock. It is used to protect shared data from being accessed by multiple goroutines at the same time.

func main () {
	fmt.Println("Hello Go, Learning Concurrency and Race Condition")
	// var wg sync.WaitGroup
	// wg.Add(2) // Add 2 to the wait group counter

	// ch := make(chan bool) // Unbuffered channel
	ch := make(chan bool, 2) // Buffered channel with a capacity of 2 works like asynchronously.

	// go increment(&wg) 
	// go increment(&wg)
	go increment(ch) 
	go increment(ch)
	// wg.Wait() // Until wg is zero
	// Wait group

	// Wait for the goroutines to complete
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
	fmt.Println("Counter:", counter)
}

// func increment (wg *sync.WaitGroup) {
func increment (ch chan bool) {
	// defer wg.Done() // wg = wg - 1
	for range 1000 {
		mu.Lock() // Lock the mutex to prevent other goroutines
		counter++
		mu.Unlock() // Unlock the mutex to allow other goroutines to access the counter
	}
	ch <- true // Send a value to the channel to indicate that the goroutine has completed its work. This allows the main goroutine to wait for the new goroutine to finish before exiting the program.
}
