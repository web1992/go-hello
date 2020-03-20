package reflecttest

import (
	"fmt"
	"reflect"
)

// Person test
type Person struct {
}

// Reflect test
func Reflect() {

	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(3.3))
	fmt.Println(reflect.TypeOf("asd"))

	person := new(Person)

	fmt.Println(reflect.TypeOf(person))
}
