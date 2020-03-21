package iotatest

// Month int
type Month int

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

// IotaTest iota test
func IotaTest() {
	println("December is", December)
	println("January is", January)
}
