package main

import (
	"fmt"
	"time"
)

func main() {
	f := fmt.Println

	t := time.Now()
	f(t.Format(time.RFC3339))
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	f(t1)

	f(t.Format("3:04PM"))
	f(t.Format("Mon Jan _2 15:04:05 2006"))
	f(t.Format("2006-01-02 15:04:05.000"))
	f(t.Format("01/02/2006 15:04:05.000"))

	t2, _ := time.Parse("2006-01-02 15:04:05.000", "2016-12-13 14:20:31.123")
	f(t2)
}
