package errorstest

import (
	"errors"
	"fmt"
)

// ErrorTest
func ErrorTest() {
	err := errors.New("I am error")
	fmt.Println(err)
}
