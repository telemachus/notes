package main

import "fmt"

func fibMemo(n int) int {
	fibs := make(map[int]int)
	fibs[0] = 0
	fibs[1] = 1

	var memFib func(n int, fibs map[int]int) int
	memFib = func(n int, fibs map[int]int) int {
		switch {
		case n < 0:
			panic("Don’t feed fib a negative number!")
		default:
			if f, ok := fibs[n]; ok {
				return f
			}
			f := memFib(n-1, fibs) + memFib(n-2, fibs)
			fibs[n] = f
			return f
		}
	}

	return memFib(n, fibs)

}

func main() {
	fmt.Printf("fib(50) = %d\n", fib(50))
}
