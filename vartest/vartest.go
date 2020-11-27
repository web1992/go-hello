package vartest

import "fmt"

var (
	// GlobalVar 全局变量
	GlobalVar = "I am global var."
)

type User struct {
	Name  string
	Level int32
}

// VarTest test var
func VarTest() {
	s := 1
	fmt.Println(s)

	fmt.Println("Global Var is:[", GlobalVar, "]")

	var name string
	name = "Lucy"
	fmt.Println("name is:", name)

	var c = "1312"
	fmt.Println("c is:", c)

	var c2 string = "1312"
	fmt.Println("c2 is:", c2)
}
