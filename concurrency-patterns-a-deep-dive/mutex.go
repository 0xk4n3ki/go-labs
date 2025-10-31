package main

import (
	"fmt"
	"sync"
	// "math/rand"
	// "time"
)

var (
	counter int
	mu sync.Mutex
)

func main() {
	numWorkers := 3
	numJobs := 10

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All jobs processed. Final counter", counter)
}

func heavyWork(id int) int {
	// time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	return id*2
}

func worker(id int, jobs <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result := heavyWork(job)

		mu.Lock()
		counter++
		fmt.Printf("Worker %d processed job %d, counter = %d\n", id, job, counter)
		mu.Unlock()

		// time.Sleep(time.Millisecond * 10)
		_ = result
	}
}