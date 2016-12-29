package main

import (
	"fmt"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func main() {
	var strs = []string{"a", "b", "c", "d"}

	fmt.Println(Index(strs, "b"))
	fmt.Println(Include(strs, "e"))
	fmt.Println(All(strs, func(v string) bool {
		return v == "a"
	}))
}
