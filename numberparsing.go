package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	f, _ = strconv.ParseFloat("1.234", 16)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	i, _ = strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	i, _ = strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(i)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
