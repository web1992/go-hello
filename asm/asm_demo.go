package main

import "fmt"

// go tool compile -S file.go
func main() {
	i := 1
	b := i + 1

	fmt.Println(b)
}
