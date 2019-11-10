package vartest

import "fmt"

var (
	// GlobaVar 全局变量
	GlobaVar = "I am global var."
)

// VarTest test var
func VarTest() {
	s := 1
	fmt.Println(s)
}
