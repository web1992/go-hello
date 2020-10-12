package newtest

import "fmt"

func newInt() {
	i := int(1)
	fmt.Println(i)
}

func newFloat() {

	f := new(float64)
	fmt.Printf("f is %f \n", *f)
	*f = 1.1
	fmt.Printf("f is %f \n", *f)
}

type user struct {
	firstName string
	lastName  string
}

func (u user) hello() {
	fmt.Printf("Hello %s %s \n", u.firstName, u.lastName)
}

func newStruct() {

	//u := new(user)
	u := user{"xx", "aa"}
	u.hello()
	u.firstName = "Lucy"
	u.lastName = "001"
	u.hello()

}
