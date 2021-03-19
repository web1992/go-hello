package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring6(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func main() {

	t := time.After(3 * time.Second)

	c := boring6("Tom")

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-t:
			{
				fmt.Println("out time")
				return
			}
		}
	}
}
