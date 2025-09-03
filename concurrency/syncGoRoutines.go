package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}

func worker(id int) {
	defer wg.Done()

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)

	fmt.Printf("worker %d done\n", id)
}