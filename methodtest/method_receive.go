package methodtest

import (
	"fmt"
	"reflect"
)

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

type shape interface {
	area()
	double()
}

type rectangle struct {
	w float64
	h float64
}

func (r rectangle) area() {
	fmt.Printf("area is %f \n", r.w*r.h)
}

func (r *rectangle) double() {
	r.w = r.w * 2
	r.h = r.h * 2
}

func TestReceive() {

	var r = rectangle{2, 3}

	r.area()
	r.double()
	r.area()
}

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func Test1() {
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
