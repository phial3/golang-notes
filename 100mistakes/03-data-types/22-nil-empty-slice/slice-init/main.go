package main

import (
	"fmt"
)

func main() {
	var s []string
	log(1, s)

	s = []string(nil)
	log(2, s)

	s = []string{}
	log(3, s)

	s = make([]string, 0)
	log(4, s)
}

func log(i int, s []string) {
	fmt.Printf("%d: cap=%d\tlen=%d\tnil=%t\n", i, cap(s), len(s), s == nil)
}
