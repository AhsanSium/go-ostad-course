package main

import (
	"fmt"
	"net/http"
	"time"
)

func main () { // Main Go Routine
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://linkedin.com",
		"http://x.com",
		"http://github.com",
	}
	ch := make(chan string) // Channel created to communicate between goroutines. Channels are a way for goroutines to communicate with each other and synchronize their execution. They provide a way to send and receive values between goroutines, allowing them to coordinate their work and share data safely.
	for _, link := range links {
		go checkLink(link, ch)
	}
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// for {
	// 	go checkLink(<-ch, ch) // New goroutine created to run checkLink function concurrently. Managed by the Go runtime scheduler, which is responsible for scheduling and executing goroutines. The Go runtime scheduler uses a work-stealing algorithm to distribute goroutines across available threads, allowing for efficient use of system resources.
	// }

	// Using Range, Range is a built-in function in Go that allows you to iterate over elements in a variety of data structures, including arrays, slices, maps, and channels. When used with channels, range allows you to receive values from the channel until it is closed. This is useful for processing a stream of data or for coordinating the execution of multiple goroutines.
	for link := range ch {
		// go checkLink(link, ch) // New goroutine created to run checkLink function concurrently. Managed by the Go runtime scheduler, which is responsible for scheduling and executing goroutines. The Go runtime scheduler uses a work-stealing algorithm to distribute goroutines across available threads, allowing for efficient use of system resources.

		go func (link string, ch chan string) {
			time.Sleep(5 * time.Second) // Sleep for 5 seconds to simulate some work being done in the goroutine. This is a simple way to demonstrate the concurrent execution of goroutines, but in a real-world application, you would typically perform some meaningful work in the goroutine.
			// time.Tick(5 * time.Second) // Sleep for 5 seconds to simulate some work being done in the goroutine. This is a simple way to demonstrate the concurrent execution of goroutines, but in a real-world application, you would typically perform some meaningful work in the goroutine.
			fmt.Println("5 Second Gap -------------: ")
			checkLink(link, ch) // New goroutine created to run checkLink function concurrently. Managed by the Go runtime scheduler, which is responsible for scheduling and executing goroutines. The Go runtime scheduler uses a work-stealing algorithm to distribute goroutines across available threads, allowing for efficient use of system resources.
		}(link, ch)
	}

}

func checkLink (Link string, ch chan string) {
	_, err := http.Get(Link)
	if err != nil {
		fmt.Println(Link+" is Down (X):", err)
		ch <- Link
	} else {
		fmt.Println(Link+" is UP (✓)")
		ch <- Link
	}
}
