package iotatest

import "fmt"

// Month int
type Month int
type WeekDay int

const PI = 3.14
const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

const (
	Monday WeekDay = 1 + iota
	Tuesday
)

// IotaTest iota test
func IotaTest() {
	println("December is", December)
	println("January is", January)
	println("Monday is", Monday)
	println("Tuesday is", Tuesday)
	println("PI is", PI)

	fmt.Printf("PI is %T \n", PI)
	fmt.Printf("PI is %F \n", PI)

}
