package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Advanced Channel
var MAX_CHICKEN_PRICE float32 = 5

func main() {
	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("\nFound a deal on chicken at %s", <-chickenChannel)
}

// Channels + Goroutines
/*func main() {
	var c = make(chan int, 5)
	go process(c)
	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	// using defer to tell the function to close the channel right before it exits, vs making sure to add it at the last line
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting the process")
	// close(c)
}*/
