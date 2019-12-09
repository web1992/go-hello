package pointertest

import (
	"fmt"
)

// Pointertest test pointer
func Pointertest() {

	a := 1
	var ap *int = &a

	fmt.Println("ap is", ap)

	*ap++

	fmt.Println("ap after add is", *ap)

}
