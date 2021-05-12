package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"https://www.apple.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.microsoft.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			bs, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize
	for range urls {
		result := <-results
		fmt.Println("Current page and size:", result.URL, result.Size)
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL, biggest.Size)
}
