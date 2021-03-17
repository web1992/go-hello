package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring7(msg string, quit chan string) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("%s %d", msg, i):
				// noop
			case <-quit:
				quit <- "See you"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func main() {

	quit := make(chan string)
	c := boring7("Tom", quit)

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	quit <- "Bey"
	fmt.Println("Tom say:", <-quit)

}
