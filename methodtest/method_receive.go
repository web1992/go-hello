package methodtest

import "fmt"

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
