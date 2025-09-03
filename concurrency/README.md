- [https://bwoff.medium.com/the-comprehensive-guide-to-concurrency-in-golang-aaa99f8bccf6](https://bwoff.medium.com/the-comprehensive-guide-to-concurrency-in-golang-aaa99f8bccf6)

### Core Idea
- Concurrency is the bigger idea → dealing with many tasks at once.
- Parallelism is one way to achieve concurrency → by literally running tasks at the same instant.

Think of it like this:
- Concurrency = umbrella term → "I’ll make progress on multiple tasks at once."
- Parallelism = special case → "I’ll make progress on multiple tasks at once because I have multiple workers."

### Which one is better?

It depends on the problem:
- Concurrency is better when tasks involve waiting (I/O heavy).
    - Example: handling thousands of web requests in a server. While one request waits for a database, you can handle others.
    - Saves CPU from sitting idle.
- Parallelism is better when tasks involve pure computation (CPU heavy).
    - Example: crunching big numbers, training machine learning models, image/video processing.
    - Splits the work across CPU cores → finishes faster.

In practice:
- Modern systems often combine both.
    - A web server handles requests concurrently.
    - Inside each request, a heavy computation might run in parallel.

### Goroutines

- Goroutines are like super-lightweight threads in Go.

    They let you run functions independently without blocking others.
- By default, when you start Goroutines:
    - They run concurrently (taking turns on one CPU core).
    - They are not truly parallel unless Go decides to spread them across multiple cores.
- GOMAXPROCS is a Go setting that tells the runtime:
    - “Use this many CPU cores at the same time.”
    - Example: If you set GOMAXPROCS = 4 on a 4-core CPU, Goroutines can run in parallel.

So:
- Goroutines = concurrency (task juggling).
- Goroutines + multiple cores (via GOMAXPROCS) = parallelism (real teamwork).

#### Advantages of Go’s Concurrency Model

1. Resource efficiency
    - Normal OS threads are “heavy” (need MBs of memory each).
    - Goroutines are “tiny” (start with KBs and grow/shrink).
    - You can have millions of Goroutines without killing your system.
2. Easy synchronization (Channels)
    - Instead of messy locks, Go gives channels to safely pass data.
    - This reduces bugs like deadlocks or race conditions.
3. Strong standard library
    - Go has built-in packages (sync, sync/atomic) for coordination.
    - No need for third-party libraries for most common concurrency needs.

#### Disadvantages
1. Concurrency ≠ Parallelism
    - Just because you use Goroutines doesn’t mean your program runs faster.
    - Unless you use multiple cores, it’s concurrent but not parallel.
2. Shared state & data races
    - If multiple Goroutines touch the same variable without coordination → race condition.
    - Go’s philosophy is “don’t share memory, communicate over channels”, but in practice people still share variables → problems.
3. Debugging complexity
    - Concurrency makes program execution non-deterministic (order changes every run).
    - Debugging such issues is harder.
    - Go helps with tools like `race detector` and `pprof`, but they take practice.

### Simple Analogy (Go version)
- Goroutines = tiny workers.
- Channels = pipes they use to pass messages safely.
- GOMAXPROCS = how many kitchens (CPU cores) you let them use.
    - 1 kitchen → they take turns (concurrency).
    - Many kitchens → they cook side by side (parallelism).


```go
package main

import (
	"fmt"
	"time"
)

func printMessage() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go printMessage()
	fmt.Println("Hello from main function")
	time.Sleep(time.Millisecond)
}
```


### Channels

1. Channels synchronize Goroutines
    - Without channels, Goroutines run independently, and you don’t know when they finish.
    - With channels, you can wait for results safely.
2. Send & Receive
    - `message <- "Hello"` → send into channel.
    - `msg := <-message` → receive from channel.
3. Blocking behavior
    - Sending waits until someone receives.
    - Receiving waits until someone sends.
    - This is why channels are great for coordination.

In short:
- Goroutines = workers.
- Channels = walkie-talkies to coordinate and exchange results.
- `chan int` → send + receive (default).
- `chan<- int` → send-only.
- `<-chan int` → receive-only.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	go printMessage(message)

	fmt.Println("Hello from main function")
	fmt.Println(<-message)
}

func printMessage(message chan string) {
	time.Sleep(time.Second*2)
	message <- "hello from goroutine"
}
```


### Goroutines: Key Points
- What they are:
    - Goroutines are lightweight functions that run concurrently.
    - They are managed by the Go runtime, not the OS.
- Scheduling (m:n model):
    - Many Goroutines (m) are multiplexed onto fewer OS threads (n).
    - Go’s runtime scheduler handles this efficiently in user space.
- Memory efficiency:
    - Start with a very small stack (~2 KB vs. ~1–2 MB in traditional threads).
    - Stack grows and shrinks dynamically as needed.
    - This allows thousands or even millions of Goroutines to run without exhausting memory.
- Blocking operations:
    - In traditional threads → blocking (like I/O) puts the thread to sleep, causing expensive context switches.
    - In Go → if one Goroutine blocks, the runtime automatically moves others to runnable threads → keeps work going.
- Why they’re powerful:
    - Lightweight, scalable, and efficient.
    - Enable high concurrency with low resource usage.
    - Simpler to use compared to traditional threads.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println("Results:", <- results)
	}
}

func worker(id int, jobs <- chan int, results chan  <- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j*2
	}
}
```


### WaitGroup Basics
- A WaitGroup lets you wait until all Goroutines finish.
- You don’t create it with new; just use var wg sync.WaitGroup.
- Three important methods:
    - `wg.Add(n)` → tell the WaitGroup how many Goroutines you’ll wait for.
    - `wg.Done()` → each Goroutine calls this when it’s finished (decreases the counter by 1).
    - `wg.Wait()` → blocks the main Goroutine until the counter is 0.

```go
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
```