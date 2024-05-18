package main

import (
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("?puS puS puS"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
