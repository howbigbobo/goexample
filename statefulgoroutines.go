package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				fmt.Println("reading...")
				read.resp <- state[read.key]
				fmt.Println("readed")
			case write := <-writes:
				fmt.Println("writing...")
				state[write.key] = write.val
				fmt.Println("writed")
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 1; r++ {
		go func() {
			read := &readOp{
				key:  rand.Intn(5),
				resp: make(chan int)}
			fmt.Println("befor read")
			reads <- read
			fmt.Println("after read")
			<-read.resp
			fmt.Println("finish read")
			atomic.AddInt64(&ops, 1)
		}()
	}

	for w := 0; w < 1; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				fmt.Println("before write")
				writes <- write
				fmt.Println("after write")
				<-write.resp
				fmt.Println("write finish")
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Millisecond * 1)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
