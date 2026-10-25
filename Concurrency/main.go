/*
Concurrency is the ability of a program to make progress on more than one task at the same time. In Go, concurrency is achieved through goroutines and channels.

Parallelism is the ability of a program to make progress on more than one task at the same time by utilizing multiple processors or cores. In Go, parallelism is achieved through goroutines and channels, but it also requires the use of multiple processors or cores.
*/
package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println("Hello Go") // Main goroutine, Main thread. The main function runs in the main goroutine, which is the initial goroutine created when a Go program starts. The main goroutine is responsible for executing the main function and any other functions that are called from it.
	ch := make(chan string) // Channel created to communicate between goroutines. Channels are a way for goroutines to communicate with each other and synchronize their execution. They provide a way to send and receive values between goroutines, allowing them to coordinate their work and share data safely.
	go printMessage("World", ch) // New goroutine created to run printMessage function concurrently. Managed by the Go runtime scheduler, which is responsible for scheduling and executing goroutines. The Go runtime scheduler uses a work-stealing algorithm to distribute goroutines across available threads, allowing for efficient use of system resources.
	time.Sleep(1 * time.Microsecond) // the main goroutine sleeps for 1 microsecond to allow the new goroutine to complete its execution before the program exits. This is a simple way to synchronize the execution of goroutines, but it is not a reliable method for ensuring that all goroutines have completed their work before the program exits. In a real-world application, you would typically use channels or other synchronization mechanisms to coordinate the execution of goroutines.

	// Sleep is not a reliable way to synchronize goroutines, as it can lead to race conditions and other issues. Instead, you should use channels or other synchronization mechanisms to coordinate the execution of goroutines. For example, you could use a WaitGroup to wait for all goroutines to complete before exiting the program. We need channels to communicate between goroutines and synchronize their execution. Channels provide a way for goroutines to send and receive values, allowing them to coordinate their work and share data safely. By using channels, we can avoid race conditions and other issues that can arise when multiple goroutines access shared data concurrently.
	message := <-ch
	fmt.Println("Message from goroutine: ", message) // receive the message from the channel and print it to the console. This allows the main goroutine to wait for the new goroutine to finish before exiting the program. 
}

func printMessage (message string, ch chan string) {
	fmt.Println("Hello: ",message)
	for i := 0; i < 5; i++ {
		fmt.Println("Printing message: ", message, " - ", i)
		// time.Sleep(1 * time.Second) // Sleep for 1 second to simulate some work being done in the goroutine. This is a simple way to demonstrate the concurrent execution of goroutines, but in a real-world application, you would typically perform some meaningful work in the goroutine.
	}
	ch <- "complete" // send a message to the channel to indicate that the goroutine has completed its work. This allows the main goroutine to wait for the new goroutine to finish before exiting the program.
}