package interfacetest

import (
	"fmt"
)

// Car interface
type Car interface {
	run()
	start() string
}

// BenzCar struct
type BenzCar struct {
	name string
}

func (car *BenzCar) run() {
	fmt.Println(car.name + " Car  run ...")
}

func (car *BenzCar) start() string {
	return car.name + " Car  start ..."
}

// InterfaceTest test
func InterfaceTest() {

	carBenz := new(BenzCar)
	carBenz.name = "I am Benz "
	carBenz.run()
	rs := carBenz.start()

	fmt.Println("result is  ", rs)

	var car1 Car
	car1 = carBenz
	car1.run()

	fmt.Println("car1 is  ", car1)

}
