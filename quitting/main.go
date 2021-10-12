package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan string)
	c := boring("Ali", quit)
	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}
	quit <- "Bye"
	fmt.Printf("Ali %s dedi", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				fmt.Println("Clean up!")
				quit <- "See you!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
