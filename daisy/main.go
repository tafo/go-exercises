package main

import "fmt"

func main() {
	const n = 10000

	var channels [n+1]chan int

	for i:= range channels {
		channels[i] = make(chan int)
	}

	for i := 0; i < n; i++ {
		go f(channels[i], channels[i+1])
	}

	go func(c chan<- int) {
		c <- 1
	}(channels[n])

	fmt.Println(<-channels[0])
}

func f(left chan<- int, right <-chan int) {
	left <- 1 + <-right
}
