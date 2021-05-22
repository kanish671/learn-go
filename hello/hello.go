package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/user/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world!!"))
	fmt.Println(cmp.Diff("Hello, world!", "Hello, go!"))
}
