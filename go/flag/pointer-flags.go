package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlag = flag.Int("n", 0, "specify the count of n")
	var sFlag = flag.String("name", "", "what is the name?")
	var bFlag = flag.Bool("loop", false, "should we loop?")
	flag.Parse()

	if *bFlag {
		for i := 0; i < *nFlag; i++ {
			fmt.Printf("sFlag = %q\n", *sFlag)
		}
	} else {
		fmt.Printf("bFlag = %t, sFlag = %q, and nFlag = %d\n", *bFlag, *sFlag, *nFlag)
	}
}
