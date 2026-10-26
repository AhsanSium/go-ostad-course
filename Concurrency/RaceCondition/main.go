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
	var wg sync.WaitGroup
	wg.Add(2) // Add 2 to the wait group counter
	go increment(&wg) 
	go increment(&wg)
	wg.Wait() // Until wg is zero
	// Wait group

	fmt.Println("Counter:", counter)
}

func increment (wg *sync.WaitGroup) {
	defer wg.Done() // wg = wg - 1
	for range 1000 {
		mu.Lock() // Lock the mutex to prevent other goroutines
		counter++
		mu.Unlock() // Unlock the mutex to allow other goroutines to access the counter
	}
}
