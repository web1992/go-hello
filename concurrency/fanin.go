package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring4(msg string) chan string {
	ch := make(chan string)
	go func() {

		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%s,%d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}
		close(ch)
	}()
	return ch

}

func fanIn(ch1, ch2 <-chan string) <-chan string {

	ch := make(chan string)
	// for ch1
	go func() {
		for {
			v := <-ch1
			ch <- v
		}
	}()

	// for ch2
	go func() {
		for {
			ch <- <-ch2
		}
	}()

	return ch
}

func fanSimple(ch2 ...<-chan string) <-chan string {

	ch := make(chan string)

	for _, chi := range ch2 {
		// use go func in for
		go func(chv <-chan string) {
			for {
				// for end when chan close
				ch <- <-chv
			}
		}(chi) // give param
	}

	return ch

}

// merge tow chan to one chan
func main() {

	ch1 := boring4("Tom")
	ch2 := boring4("Jerry")

	//c := fanIn(ch1, ch2)
	//for i := 1; i < 10; i++ {
	//	fmt.Println(<-c)
	//}

	c2 := fanSimple(ch1, ch2)
	for i := 1; i < 10; i++ {
		fmt.Println(<-c2)
	}
	fmt.Println("You're both boring. I'm leaving")
}
