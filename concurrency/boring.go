package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go boring("boring!")
	// for {

	// }

	fmt.Println("I am listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring. I'm leaving")
}
