package main

import "fmt"

func half(n int) (int, bool) {
	h := n / 2
	r := false
	if n%2 == 0 {
		r = true
	}
	return h, r
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
