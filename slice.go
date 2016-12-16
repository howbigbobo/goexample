package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("emp=", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c=", c)

	s = append(s, "d")
	s = append(s, "e")

	l := s[2:5]
	fmt.Println("l1=", l)

	l = s[:5]
	fmt.Println("l2=", l)

	l[2] = "m"
	fmt.Println(s)

	c[1] = "x"
	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(l)
}
