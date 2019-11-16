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
	)
	fmt.Println(a,b,c)


}
