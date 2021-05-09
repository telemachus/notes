package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", showStuff)

}

func showStuff(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(path, info.Size())
	return nil
}
