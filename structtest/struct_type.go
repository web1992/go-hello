package structtest

import "fmt"

// part
// declare type at package level
type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func StructType() {

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323

	var bolts part

	bolts.description = "Hex bolts"
	bolts.count = 24

	fmt.Println("porsche is", porsche)
	fmt.Println("bolts is", bolts)
	// change count
	ChangeCount(&bolts)
	fmt.Println("bolts is", bolts)

	i := 0
	fmt.Println("i is ", i)
	ChangeInt(&i)
	fmt.Println("i is ", i)
}

// change struct type not need *
func ChangeCount(p *part) {
	(*p).count = 68
	// 下面是上面的简写
	//p.count = 68
}

// change int type need *
func ChangeInt(i *int) {
	*i = 1
}
