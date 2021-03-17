package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func boring5(msg string) chan Message {
	ch := make(chan Message)
	waitFotIt := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitFotIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			// wait for do waitFotItc <-true
			<-waitFotIt
		}
	}()

	return ch
}

func fanIn5(ch ...chan Message) <-chan Message {
	chR := make(chan Message)

	for i := range ch {
		input := ch[i]
		go func() {
			for {
				chR <- <-input
			}
		}()
	}
	return chR
}

func main() {
	// merge
	c := fanIn5(boring5("Tom"), boring5("Jerry"))

	for i := 0; i < 5; i++ {
		m1 := <-c
		fmt.Println(m1.str)
		m2 := <-c
		fmt.Println(m2.str)
		m1.wait <- true
		m2.wait <- true
	}

	fmt.Println("You're both boring. I'm leaving")
}
