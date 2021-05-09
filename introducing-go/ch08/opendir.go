package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Open("ch07")
	if err != nil {
		log.Fatal("cannot open the directory 'ch07'")
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal("cannot read the directory 'ch07'")
	}
	for _, fi := range files {
		fmt.Println(fi.Name(), fi.Size())
	}
}
