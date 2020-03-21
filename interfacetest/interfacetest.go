package interfacetest

import (
	"fmt"
)

// EmptyInterface interface
type EmptyInterface interface{}

// Car interface
type Car interface {
	run()
	start() string
	bing()
}

// BenzCar struct
// 奔驰
type BenzCar struct {
	name string
}

func (car *BenzCar) run() {
	fmt.Println(car.name + " Car  run ...")
}

func (car *BenzCar) start() string {
	return car.name + " Car  start ..."
}

func (car *BenzCar) bing() {
	fmt.Println("bing is from ", car.name)
}

// InterfaceTest test
func InterfaceTest() {

	carBenz1 := new(BenzCar)
	carBenz1.name = "xxx xxx xxx"
	carBenz1.run()
	carBenz1.bing()

	carBenz := new(BenzCar)
	carBenz.name = "I am Benz "
	carBenz.run()
	rs := carBenz.start()

	fmt.Println("result is  ", rs)

	var car1 Car
	car1 = carBenz
	car1.run()

	fmt.Println("car1 is  ", car1)

	println(car1)
	println(carBenz)

	var ins2 Car
	b2 := BenzCar{"Benz 2"}
	b2.run()
	fmt.Println(ins2)
	// below is error
	// to solve this you can use b2 := BenzCar{"Benz 2"}
	// ins2 = b2
	// ins2.run()

	var ei EmptyInterface
	ei = b2
	fmt.Println("ei ", ei)

}
