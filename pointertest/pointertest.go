package pointertest

import (
	"fmt"
)

// PointerTest test pointer
func PointerTest() {

	a := 1
	//var ap *int = &a
	ap := &a

	fmt.Println("a pointer is", ap)

	*ap++

	fmt.Println("*ap after add is", *ap)

	fmt.Println("======================")
	v := 2
	fmt.Println("v pointer is", &v)
	*(&v)++
	fmt.Println("*v is after add", v)

}
