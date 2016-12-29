package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("v: ", i)

	zeroptr(&i)
	fmt.Println("p: ", i)

	fmt.Println(" ptr: ", &i)
}
