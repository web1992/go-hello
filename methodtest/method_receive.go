package methodtest

import (
	"fmt"
	"reflect"
)

// https://mp.weixin.qq.com/s/XKirIaGmyBAwpFKo9Yekxw

func printArea(s shape) {
	s.area()
}

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
