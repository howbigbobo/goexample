package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, e := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	fmt.Println(e)

	r, _ := regexp.Compile("p([a-z]+)ch")

}
