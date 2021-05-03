package main

import "testing"

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

func fibBad(n int) int {
	switch {
	case n < 0:
		panic("Don’t feed fib a negative number!")
	case n <= 1:
		return n
	default:
		return fibBad(n-1) + fibBad(n-2)
	}
}

func fibIter(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func BenchmarkFibMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibMemo(50)
	}
}

func BenchmarkFibBad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibBad(50)
	}
}

func BenchmarkFibIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibIter(50)
	}
}
