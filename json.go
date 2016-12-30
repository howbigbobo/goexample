package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page2"`
	Fruits []string `json:"fruits2"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal([]string{"a", "b"})
	fmt.Println(string(intB))

	mapB, _ := json.Marshal(map[string]int{"a": 1, "b": 2})
	fmt.Println(string(mapB))

	res1 := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res1B, _ := json.Marshal(res1)
	fmt.Println(string(res1B))

	res2 := &Response2{
		Page:   2,
		Fruits: []string{"a", "b"}}
	res2B, _ := json.Marshal(res2)
	fmt.Println(string(res2B))

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
