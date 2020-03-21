package channeltest

import (
	"fmt"
	"time"
)

type empty interface{}

// Car struct
type Car struct {
	name   string
	wheels int
}

// ChannelTest test
func ChannelTest() {

	ch := make(chan int)

	fmt.Println("ch", ch)
	fmt.Println("ch pointer", &ch)

	go func() {
		ch <- 1
		ch <- 2
	}()

	val, ok := <-ch
	fmt.Println("val ", val, "ok", ok)

	empty := <-ch
	fmt.Println("empty is", empty)

	close(ch)
	go func() {
		// panic: send on closed channel
		//ch <- 3
	}()

	time.Sleep(5)
	v3, ok := <-ch
	v4, ok := <-ch
	fmt.Println("v3", v3, ok)
	fmt.Println("v4", v4, ok)

	// buffer channel
	bufCh := make(chan int, 5)
	fmt.Println("bufCh", bufCh)

	bufCh <- 1 // 1
	bufCh <- 1 // 2
	bufCh <- 1 // 3
	bufCh <- 1 // 4
	bufCh <- 1 // 5
	// 6
	// bufCh <- 1

	// time.Sleep(1e9)

	fmt.Println(time.Now())
	fmt.Println(time.Second)
}

// RequestTest test
func RequestTest() {

	carCh := make(chan *Car, 2)
	start := make(chan int)

	go forChannelTest(carCh, start)

	c := &Car{"BenZ Car", 4}
	c2 := &Car{"BenZ Car2", 8}

	carCh <- c
	carCh <- c2

	// time.Sleep(time.Second * 1)

	start <- 1
}

func forChannelTest(queue chan *Car, start chan int) {

	<-start

	for c := range queue {
		fmt.Println("car is", c)
	}
}
