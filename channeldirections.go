package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	fmt.Println("done1")

	time.Sleep(time.Second)
	fmt.Println("done2")
	<-done
	fmt.Println("done3")

	pings := make(chan string, 1)
	pongs := make(chan string, 2)
	ping(pings, "pass")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
