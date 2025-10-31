package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	for i := 0; i < 3; i++ {
		go once.Do(initialize)
		go once.Do(hello)
		go once.Do(suraj)
	}
	time.Sleep(100 * time.Millisecond)
}

func initialize() {
	fmt.Println("Initialization complete")
}

func hello() {
	fmt.Println("Hello World")
}

func suraj() {
	fmt.Println("Hello Suraj")
}