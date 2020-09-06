package main

import (
	"fmt"
	"os"
)

func main() {
	// print args
	fmt.Println("Hello world!")
	fmt.Println(os.Args[1:])

}
