package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ali := boring("Ali")

	timeout := time.After(5 * time.Second)

	for {
		select {
		case str := <-ali:
			fmt.Println(str)
		case a := <-timeout:
			fmt.Println("That's enough = ", a)
			return
		}
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
