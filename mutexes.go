package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(15)
				mutex.Lock()
				fmt.Print("r ")
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(15)
				val := rand.Intn(100)
				mutex.Lock()
				fmt.Print("w ")
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Millisecond * 10)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()
}
