package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlag int
	flag.IntVar(&nFlag, "n", 0, "specify the count of n")
	var sFlag string
	flag.StringVar(&sFlag, "name", "", "what is the name?")
	var bFlag bool
	flag.BoolVar(&bFlag, "loop", false, "should we loop?")
	flag.Parse()

	if bFlag {
		for i := 0; i < nFlag; i++ {
			fmt.Printf("sFlag = %q\n", sFlag)
		}
	} else {
		fmt.Printf("bFlag = %t, sFlag = %q, and nFlag = %d\n", bFlag, sFlag, nFlag)
	}
}

