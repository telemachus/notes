package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	s := "I 想 want 我to eat!吃"
	r := strings.NewReader(s)
	data := make([]byte, 1056)

	w := bufio.NewWriter(os.Stdout)
	for {
		c, err := r.Read(data)
		if err != nil {
			break
		}
		if c == 0 {
			break
		}
		// fmt.Println(c, data[:c], string(data[:c]))
		w.Write(data[:c])
	}
	w.Write([]byte("\n"))
	w.Flush()
}
