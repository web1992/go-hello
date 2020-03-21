package structtest

import (
	"fmt"
)

// StructTest test
func StructTest() {

	var p person

	p.name = "Lucy"
	p.age = 18

	fmt.Println("p is", p)
	fmt.Println("p.name is", p.name)

	p2 := person{
		name: "man fen",
		age:  3,
	}

	fmt.Println("p2 is", p2)

	p3 := new(person)
	p3.name = "rou rou"
	p3.age = 3

	fmt.Println("p3 is", p3)

	p4 := &person{}
	p4.name = "Li Li"
	p4.age = 18

	fmt.Println("p4 is", p4)

	var p5 *person
	//p5 is <nil>
	fmt.Println("p5 is", p5)
	p5 = p4
	fmt.Println("after ptr p4 is", p4)

	var c C
	c.a = "asda"
	fmt.Println("C is", c)
}

type person struct {
	name string "Name"
	age  int    "Age"
}

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

type A struct {
	a int
	b int
}

type B struct {
	b float32
	c string
	d string
}

type C struct {
	A
	B
	a string
	c string
}
