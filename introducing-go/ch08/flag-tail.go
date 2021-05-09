package main

import (
	"bufio"
	"container/ring"
	"flag"
	"fmt"
	"os"
)

const n int = 10

func printRing(line interface{}) {
	if line == nil {
		return
	}
	fmt.Println(line.(string))
}

func main() {
	n := flag.Int("n", 10, "number of lines to show")
	flag.Parse()

	r := ring.New(*n)
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
