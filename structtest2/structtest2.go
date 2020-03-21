package structtest2

import "fmt"

// StructTest2 test
func StructTest2() {
	bmHorse := &Horse{
		Animal: &Animal{ // 注意此行
			name:   "Mouse",
			weight: 60,
		},
		speak: "Zi zi zi",
	}
	bmHorse.hello()
	fmt.Println(bmHorse.name)
	fmt.Println(bmHorse.speak)
}

// Animal struct
type Animal struct {
	name   string
	weight int
}

// Horse struct
type Horse struct {
	*Animal // 注意此行
	speak   string
}

func (a *Animal) hello() {
	fmt.Println(a.name)
	fmt.Println(a.weight)
	//fmt.Println(a.speak)
}
