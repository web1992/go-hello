package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring2(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
	}
}

func main() {
	ch := make(chan string)
	go boring2("boring!", ch)

	for i := 0; i < 5; i++ {

		fmt.Println("you say", <-ch)
	}
	fmt.Println("You're boring. I'm leaving")

}
