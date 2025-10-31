package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string {"https://example.com", "https://fb.com", "https://google.com", "https://github.com"}
	ch := make(chan string)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go scrapeURL(url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for status := range ch {
		fmt.Println("status:", status)
	}
}

func scrapeURL(url string, ch chan <- string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer response.Body.Close()
	ch <- response.Status
}