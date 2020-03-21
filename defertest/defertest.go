package defertest

import (
	"fmt"
)

var x int = 10

// DeferTest test
func DeferTest() {

	fmt.Println("A run")
	defer afterFuncA("Test 01")
	fmt.Println("B run")

	// defer afterFunc("Test 02")
	// a := 1 - 1
	// fmt.Println(1 / a)

	x = 20
	defer afterFuncB(x) // 输出20
	x = 30
	println("after x =30 ", x) // 输出30

	defer func(a int) {
		println("inner func 匿名函数:", a) // 输出30
	}(x)

	if 1 == 1 {
		//panic("error rrrrrrrrrrrrrr.")
	}

	println("I am last .")

}

func afterFuncA(s string) {
	fmt.Println("afterFuncA after run", s)
}

func afterFuncB(x int) {
	println("afterFuncB start afterFuncB:", x)
}
