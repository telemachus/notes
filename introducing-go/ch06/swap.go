package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	x, y := 1, 2
	fmt.Printf("x = %d, and y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("x = %d, and y = %d\n", x, y)
}
