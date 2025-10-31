package main

import "fmt"

func main() {
	data := []int {1, 2, 3, 4, 5}
	input := make(chan int, len(data))

	for _, i := range data {
		input <- i
	}
	close(input)

	doubleOutput := make(chan int)
	go func() {
		defer close(doubleOutput)
		for num := range input {
			doubleOutput <- num*2
		}
	}()

	squareOutput := make(chan int)
	go func() {
		defer close(squareOutput)
		for num := range doubleOutput {
			squareOutput <- num * num
		}
	}()

	for result := range squareOutput {
		fmt.Println(result)
	}
}