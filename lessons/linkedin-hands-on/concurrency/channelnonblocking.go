package main

import (
	"fmt"
	"time"
)

func main() {
	// declare a signal channel to read from
	timeChan1 := time.After(200 * time.Millisecond)
	timeChan2 := time.After(400 * time.Millisecond)

	// use a default case in select to prevent blocking
	for {
		select {
		case <-timeChan1:
			fmt.Println("received from timeChan1")
			return
		case <-timeChan2:
			fmt.Println("received from timeChan2")
			return
		default:
			fmt.Println("no signal received")
			time.Sleep(250 * time.Millisecond)
		}
	}

}
