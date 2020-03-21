package consttest

import (
	"fmt"
)

// ConstTest test
func ConstTest() {

	const str = "dog const"
	fmt.Println(str)

	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)

	const x = "1234"
	fmt.Println("x is", x)

	//x = "a"
	// error
}
