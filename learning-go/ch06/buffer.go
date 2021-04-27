package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	count := 0
	s := "I 想 want 我to eat!吃\n"
	r := strings.NewReader(s)
	// data := make([]byte, 1)

	// w := bufio.NewWriter(os.Stdout)
	for {
		c, err := io.Copy(os.Stdout, r)
		// c, err := r.Read(data)
		if err != nil {
			break
		}
		count++
		fmt.Println("buffer read #", count)
		if c == 0 {
			break
		}
		// fmt.Println(c, data[:c], string(data[:c]))
		// w.Write(data[:c])
	}
	// w.Write([]byte("\n"))
	// w.Flush()
}
