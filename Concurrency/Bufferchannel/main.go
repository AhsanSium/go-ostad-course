package main

import (
	"fmt"
	"time"
)

func main () {
	ch := make(chan int, 2)

	go producer(ch)
	go consumer(ch)
	time.Sleep(5 * time.Second)
}

func producer (ch chan<- int) { // Only send to channel
	for i := 1; i <= 5; i++ {
		fmt.Println("Producing: ===>>>", i)
		ch <- i
		// fmt.Println("Produced:", i)
		time.Sleep(500 * time.Microsecond)
	}
	fmt.Println("Producing Done")
	close(ch) // Close the channel to indicate that no more values will be sent.
}

func consumer (ch <-chan int) { // Only receive from channel
	for val := range ch {
		fmt.Println("Consumed: <<<===", val)
		time.Sleep(2 * time.Second)
	}
}