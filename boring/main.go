package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	go boring("boring!")
	fmt.Println("I am listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You are boring. I'm leaving.")
}

func boring(msg string) {
	for i:=0; ;i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}