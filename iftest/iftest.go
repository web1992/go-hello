package iftest

import (
	"fmt"
)

// IfTest test
func IfTest() {

	x := 11

	fmt.Println(x)

	if x > 10 {
		fmt.Println("x is big then 10", x)
	}

	a := 1
	if a > 0 {
		fmt.Println("a is ", a)
	}

	str := "abcdef"

	fmt.Println("index and val ")
	// index and val
	for index, s := range str {
		fmt.Print(index, "-", s, "\t")
	}
	fmt.Println()

	fmt.Println("only index ")
	// only index
	for index1 := range str {
		fmt.Print(index1, "\t")
	}

	fmt.Println()
	fmt.Println("only val ")
	// juset print val
	for _, v := range str {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}

// TestFor test
func TestFor() {

	str := "abcd123456"
	for index, val := range str {
		fmt.Println(index, val)
	}
}
