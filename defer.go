package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("d:/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("wrting")
	fmt.Fprintln(f, "dataaaa")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
