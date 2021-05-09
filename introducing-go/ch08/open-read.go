package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("ch08/chapter-08.md")
	if err != nil {
		log.Fatal("ouch, wet the bed trying to open 'chapter-08.md'")
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal("ouch, wet the bed trying to stat 'chapter-08.md'")
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal("ouch, wet the bed trying to read 'chapter-08.md'")
	}

	text := string(bs)
	fmt.Println(text)
}
