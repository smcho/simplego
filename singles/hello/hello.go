package main

import (
	"fmt"

	"github.com/smcho/simplego/stringutil"
)

func main() {
	s := "안녕, 난 Go~"
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", stringutil.Reverse(s))
}
