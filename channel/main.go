package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)
	go boring("boring!", ch)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-ch)
	}
	fmt.Println("You're boring. I'm leaving.")
}

func boring(msg string, ch chan string) {
	for i := 0; ; i++ {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		ch <- fmt.Sprintf("%s %d", msg, i)
	}
}
