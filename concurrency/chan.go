package main

import "fmt"

func main() {
	ch := make(chan int)

	// fatal error: all goroutines are asleep - deadlock!
	go func() { ch <- 1 }()
	fmt.Println(<-ch) // will block until write a value to chan `ch <- 1`
}
