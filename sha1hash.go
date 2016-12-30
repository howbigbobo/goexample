package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "1"

	h := sha1.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	h2 := md5.New()
	h2.Write([]byte(s))
	bs2 := h2.Sum(nil)
	fmt.Printf("%x\n", bs2)

	dst := hex.EncodeToString(bs2)
	fmt.Println(dst)
}
