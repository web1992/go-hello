package channeltest

import (
	"fmt"
	"time"
)

type empty interface{}

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
	fmt.Println("val ", val)
	fmt.Println("ok ", ok)

	empty := <-ch
	fmt.Println(empty)

	close(ch)
	go func() {
		// panic: send on closed channel
		//ch <- 3
	}()

	time.Sleep(5)
	v3, ok3 := <-ch
	v4, ok4 := <-ch
	fmt.Println("v3", v3, ok3)
	fmt.Println("v4", v4, ok4)

	// buffer channel
	bufch := make(chan int, 5)
	fmt.Println("bufch", bufch)

	bufch <- 1
	bufch <- 1
	bufch <- 1
	bufch <- 1
	bufch <- 1
	// 6
	// bufch <- 1

	// time.Sleep(1e9)

	fmt.Println(time.Now())
	fmt.Println(time.Second)
}
