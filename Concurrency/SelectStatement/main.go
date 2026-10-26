package main

import (
	"fmt"
	"time"
)

func main () {
	response := make(chan string)
	go fakeApi(response)
	select {
		case res := <-response:
			fmt.Println("Response from API: ", res)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout: No response from API")
	}
}

func fakeApi (response chan string) {
	time.Sleep(2 * time.Second)
	response <- "Received"
}
