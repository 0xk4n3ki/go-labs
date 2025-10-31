package main

import (
	"sync"
	"fmt"
)

func main() {
	input := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numWorkers := 3
	wg := sync.WaitGroup{}

	jobs := make(chan int)
	results := make(chan int, len(input))

	for i := 0; i < numWorkers; i++ {
		
		wg.Add(1)
		ch1 := make(chan int)
		go func() {
			defer wg.Done()
			defer close(ch1)
			for x := range jobs {
				ch1 <- x + 1
			}
		}()

		wg.Add(1)
		ch2 := make(chan int)
		go func() {
			wg.Done()
			defer close(ch2)
			for x := range ch1 {
				ch2 <- x + 1
			}
		}()

		wg.Add(1)
		ch3 := make(chan int)
		go func() {
			wg.Done()
			defer close(ch3)
			for x := range ch2 {
				results <- x + 1
			}
		}()
	}

	go func() {
		defer close(jobs)
		for _, i := range input {
			jobs <- i
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Printf("result: %d\n", res)
	}
}
