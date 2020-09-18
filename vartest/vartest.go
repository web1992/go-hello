package vartest

import "fmt"

var (
	// GlobalVar 全局变量
	GlobalVar = "I am global var."
)

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

}
