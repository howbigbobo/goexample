package main

import (
	"os"
)

func main() {
	panic("a problem")

	_, err := os.Create("d:/tmp/file")
	if err != nil {
		panic(err)
	}
}
