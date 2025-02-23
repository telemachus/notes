package main

import "fmt"

func main() {
	*pString := "Hello, world!"
	fmt.Println(pString, "=", *pString)
}
