package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plushPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res = plushPlus(1, 2, 3)
	fmt.Println(res)

	a, b := vals()
	fmt.Println(a, b)

	c := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(c)
}
