package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("are: ", r.area())
	fmt.Println("p:", r.perim())

	rp := &r
	fmt.Println("a:", rp.area())
	fmt.Println("p2:", r.perim())
}
