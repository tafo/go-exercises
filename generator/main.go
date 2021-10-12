package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main will call a generator function.
// A generator returns a channel.
func main() {
	rand.Seed(time.Now().UnixNano())

	ali := boring("Ali")
	veli := boring("Veli")

	for i := 0; i < 10; i++ {
		fmt.Println(<-ali)
		fmt.Println(<-veli)
	}

	fmt.Println("You're boring. I'm leaving.")
}

func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	// The channel is still open
	return ch
}
