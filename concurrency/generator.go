package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring3(msg string) <-chan string {

	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
		}

		close(ch)
	}()

	return ch

}

func main() {

	jerry := boring3("jerry")
	tom := boring3("tom")

	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-jerry)
	//	fmt.Println(<-tom)
	//}

	for m := range jerry {
		fmt.Println(m)
	}

	for m := range tom {
		fmt.Println(m)
	}
	fmt.Println("You're both boring. I'm leaving")
}
