package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Ali"), boring("Veli"))

	for i := 0; i < 5; i++ {
		msg1 := <- c
		fmt.Println(msg1.str)
		msg2 := <- c
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- false // Atadığımız boolean değerin bir önemi yok
		// Yukarıdaki satırları veya satırlardan birini silersek deadlock oluşur.
		// Çünkü main function c kanalından mesaj beklerken,
		// Bu kanala mesaj gönderen boring fonksiyonu da main'i beklemiş olur.
	}
	fmt.Println("Done")
}

func fanIn(channels ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range channels {
		channel := channels[i]
		go func() {
			for {
				c <- <-channel
			}
		}()
	}
	return c
}

type Message struct {
	str  string
	wait chan bool
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool) // Shared between all messages.
	go func() {
		for i := 0; ; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}
