package main

import (
	"fmt"
	"sync"
	// "time"
)

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup
	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
			// time.Sleep(200 * time.Millisecond)
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result:", res)
	}

	fmt.Println("All Done!")
}

func worker(id int, jobs <- chan int, results chan <- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, job)
		// time.Sleep(500*time.Millisecond)
		results <- job * 2
	}
}