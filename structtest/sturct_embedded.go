package structtest

import "fmt"

// StructEmbedded for test struct embedded
func StructEmbedded() {

	var sub Subscriber

	sub.Name = "Lucy"
	sub.Rate = 3.14
	sub.Active = false

	// Subscriber set Address info
	sub.City = "Shanghai"
	sub.Town = "People Square"
	fmt.Println("sub is ", sub)

	var emp Employee
	emp.Address.City = "ShangHai"
	emp.Address.Town = "People Square"
	emp.Name = "LiLi"
	emp.Salary = 90000

	fmt.Println("emp is ", emp)

}

// Subscriber embedded Address as anonymous field
type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address `json:"address"`
}

// Employee embedded Address
type Employee struct {
	Name    string
	Salary  float64
	Address Address
}

// Address for record address info
type Address struct {
	City string
	Town string
}
