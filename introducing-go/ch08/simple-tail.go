package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
)

const n int = 10

func printRing(line interface{}) {
	fmt.Println(line.(string))
}

func main() {
	r := ring.New(n)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		r.Value = scanner.Text()
		r = r.Next()
	}

	r.Do(printRing)
	// r.Do(func(line interface{}) {
	// 	fmt.Println(line.(string))
	// })
}
