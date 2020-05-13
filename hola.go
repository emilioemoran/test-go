package main

import (
	"fmt"

	"example.com/user/hola/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("oilimE!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}