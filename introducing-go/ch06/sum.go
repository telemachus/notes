package main

import "fmt"

func sum(ns []int) int {
	total := 0
	for _, n := range ns {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
