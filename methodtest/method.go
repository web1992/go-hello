package methodtest

import (
	"fmt"
)

type myInt int

type car struct {
	name  string
	color string
}

// 重写对象的 String 方法
func (car *car) String() string {
	return "your car name is" + car.name + " and color is " + car.color
}

func (car *car) run() {
	fmt.Println(car.name, "car run.")
}

func (car *car) start() string {

	return car.name + " started."
}

func (myInt *myInt) addOne(a int) int {

	return a + 1
}

// MethodTest test
func MethodTest() {

	c1 := new(car)
	c1.name = "Ben chi"
	c1.color = "yellow"
	s := c1.start()
	fmt.Println(s)
	c1.run()
	fmt.Println("car String is ", c1)

	var a = new(myInt)
	b := a.addOne(1)
	fmt.Println("b is ", b)
}
