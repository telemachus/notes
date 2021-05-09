package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	bs, err := os.ReadFile("ch08/chapter-08.md")
	if err != nil {
		log.Fatal("wet the bed trying to read 'ch08/chapter-08.md'")
	}
	fmt.Printf("%s", string(bs))
}
