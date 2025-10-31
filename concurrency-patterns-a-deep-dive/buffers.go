package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for d := range ch {
			fmt.Println(d)
		}
	}()


	ch <- 42
	fmt.Printf("in func\n")
	ch <- 43
	fmt.Println("after 43")
}