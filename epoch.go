package main

import "time"
import "fmt"

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(nanos)
	millis := nanos / 1000
	fmt.Println(millis)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
