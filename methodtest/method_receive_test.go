package methodtest

import (
	"fmt"
	"testing"
)

func Test11(t *testing.T) {
	Test1()
}
func TestReceive2(t *testing.T) {

	var r = rectangle{2, 3}
	// 下面加上了 shape 就和报错
	// var rs shape = rectangle{2, 3}
	// fmt.Println(typeof(rs))

	fmt.Printf("%T \n", r)
	fmt.Println(typeof(r))
	r.area()
	r.double()
	r.area()
}
