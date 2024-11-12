package main

import (
	"fmt"
)

// Channels + Goroutines
func main() {
	var c = make(chan int)
	go process(c)
	fmt.Println(<-c)
}

func process(c chan int) {
	c <- 123
}
