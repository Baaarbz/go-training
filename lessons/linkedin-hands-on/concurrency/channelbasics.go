package main

import "fmt"

// sum calculates and prints the sum of numbers
func sumInt(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}
func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch := make(chan int)
	// invoke the sum function as a goroutine with channel
	go sumInt(nums, ch)
	result := <-ch
	fmt.Printf("Result: %v\n", result)

	// Buffered channels specifying size (if size is not specified then it will be an unbuffered channel)
	ch2 := make(chan string, 2)
	ch2 <- "James"
	ch2 <- "Tom"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
