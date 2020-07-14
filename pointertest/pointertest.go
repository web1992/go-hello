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
	fmt.Printf("*ap T is %T \n", *ap)
	s := fmt.Sprintf("*ap after Type is %T", ap)
	fmt.Println(s)

	fmt.Println("======================")
	v := 2
	fmt.Println("v pointer is", &v)
	vt := fmt.Sprintf("v pointer Type is %T", &v)
	fmt.Println(vt)
	*(&v)++
	fmt.Println("*v is after add", v)

}
