package main

import "fmt"
import s "strings"

var p = fmt.Println

func main() {

	p("contains: ", s.Contains("test", "es"))
	p(s.Count("test", "t"))
	p(s.HasPrefix("test", "te"))
	p(s.HasSuffix("test", "st"))
	p(s.Index("test", "e"))
	p(s.Join([]string{"a", "b"}, "-"))
	p(s.Repeat("a", 5))
	p(s.Replace("foo", "o", "0", -1))
	p(s.Replace("foo", "o", "0", 1))
	p(s.Split("a-b-c-d-e", "-"))
	p(s.ToLower("ABC"))
	p(s.ToUpper("abc"))
	p(len("hello"))
	p("hello"[2])
}
