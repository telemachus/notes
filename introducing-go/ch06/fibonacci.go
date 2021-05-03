package main

import "fmt"

func fib(n int) int {
	switch {
	case n < 0:
		panic("Don’t feed fib a negative number!")
	case n == 0:
		return 0
	case n == 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fmt.Printf("fib(50) = %d\n", fib(50))
}
