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

	str := "abcdef"

	// index and val
	for index, s := range str {
		fmt.Println(index, s)
	}

	// only index
	for index1 := range str {
		fmt.Println(index1)
	}

	for _, v := range str {
		fmt.Println(v)
	}
}
