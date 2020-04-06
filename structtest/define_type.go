package structtest

import "fmt"

func DefineType() {

	type Salary float64

	var salary Salary
	salary = 9

	salary1 := 11

	fmt.Println("salary is ", salary)
	fmt.Println("salary1 is ", salary1)
}

type Liters float64      // 升
type Milliliters float64 // 毫升
type Gallons float64     // 加仑

func (l Liters) ToMilliliters() Milliliters {

	return Milliliters(l * 1000)
}

func (m Milliliters) ToLiters() Liters {

	return Liters(m / 1000)
}

func DefineTypeConvert() {

	liters := Liters(89.00)
	fmt.Println("liters to milliliters", liters.ToMilliliters())

	milliliters := Milliliters(6000)
	fmt.Println("milliliters to liters", milliliters.ToLiters())

	fmt.Println(milliliters + 10)
	fmt.Println(liters - 10)
}
