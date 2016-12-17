package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		fmt.Println("2")
		message <- "ping"
		fmt.Println("3")
	}()

	fmt.Println("1")
	msg := <-message
	fmt.Println(msg)

	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
