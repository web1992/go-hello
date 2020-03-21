package reflecttest

import (
	"fmt"
	"reflect"
)

// Person test
type Person struct {
	name string
}

func (p *Person) String() string {
	return p.name
}

// Reflect test
func Reflect() {

	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(3.3))
	fmt.Println(reflect.TypeOf("asd"))
	fmt.Println(reflect.TypeOf(float64(22)))
	fmt.Println(reflect.TypeOf(float32(22)))
	person := new(Person)
	person.name = "Lucy"
	fmt.Println(reflect.TypeOf(person))

	fmt.Println("p is", person)
}
