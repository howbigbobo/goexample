package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings", strs)

	ints := []int{3, 1, 5}
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)
	sort.Ints(ints)
	fmt.Println("ints", ints)
	s = sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
