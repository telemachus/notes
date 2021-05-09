package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	shaSum1 := sha1.New()
	shaSum1.Write([]byte("pick up bagels"))
	string1 := fmt.Sprintf("%x", shaSum1.Sum(nil))

	shaSum2 := sha1.New()
	shaSum2.Write([]byte("pick up bagels"))
	string2 := fmt.Sprintf("%x", shaSum2.Sum(nil))

	fmt.Printf("shaSum1 = %q\nshaSum2 = %q\nshaSum1 == shaSum2? %t\n", string1, string2, string1 == string2)
}
