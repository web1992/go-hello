package main

import "fmt"

func boring8(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 10
	leftmost := make(chan int)
	left := leftmost
	right := leftmost
	// leftmost,left,right is a same chan int
	//fmt.Println(leftmost, left, right)

	for i := 0; i < n; i++ {
		right = make(chan int)
		go boring8(left, right)
		left = right
	}

	// write a value to right chan
	// will make func boring8 block to run
	go func(r chan int) { r <- 1 }(right)

	fmt.Println(<-leftmost)
	// Output: 1001
}
