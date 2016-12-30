package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 12, 12, 123123, time.UTC)
	p(then)

	p(then.Year())
	p(then.Location())
	p(now.Location())
}
