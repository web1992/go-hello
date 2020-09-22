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

func TestReceive3(t *testing.T) {

	var r = &rectangle{2, 3}
	// 如果不加 & printArea 就会报错
	// var r = rectangle{2, 3}

	printArea(r)

	fmt.Printf("%T \n", r)
	fmt.Println(typeof(r))
	r.area()
	r.double()
	r.area()
}
func Test1(t *testing.T) {
	var c coder = &Gopher{"Go"}
	var c1 coder = &Gopher{"Go"}
	var c2 = Gopher{"Go"}
	fmt.Println("var c coder is " + typeof(c))
	fmt.Println("var c1 coder is " + typeof(c1))
	fmt.Println("var c2 coder is " + typeof(c2))
	c.code()
	c.debug()
	c2.debug()
}

func test2() {
	//var c coder = Gopher{"Go"}
	//c.code()
	//c.debug()
}
