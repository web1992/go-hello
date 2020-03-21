package functest

import (
	"errors"
	"fmt"
)

// FuncTest test
func FuncTest() {

	deferTest()

	f := func(x, y int) int {
		return x + y
	}
	funcTest("Hello func", f)

	addF := add()
	v := addF(1, 2)
	fmt.Println("addF result", v)

	sliceParamTest(1, "a", "b", "c", "d")
	// slice as param
	s := []string{"e", "f", "g", "h"}
	sliceParamTest(2, s...)

	// use func as param
	useMultiReturnValuesFunc(multiReturnValuesTest())

	// panicTest
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover")
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}

	}()

	// if panic arise， code after panic will not run
	// panicTest()
	// fmt.Println("run after panic")

	v, error := errorsTest(1, 0)
	if error != nil {
		fmt.Println("errorsTest error", v, error)
	} else {
		fmt.Println("errorsTest ok", v)
	}

	v2, error2 := errorsTest(4, 2)
	if error2 != nil {
		fmt.Println("errorsTest error ", v2, error2)
	} else {
		fmt.Println("errorsTest ok ", v2)
	}

	e := error.Error()
	fmt.Println("Error ", e)

}

// funcTest test
// func as param
func funcTest(msg string, f func(x, y int) int) {
	fmt.Println(msg, ":", f(1, 1))
}

// add func
// return func
func add() func(a, b int) int {

	f := func(a, b int) int {
		return a + b
	}
	return f
}

// sliceParamTest test
func sliceParamTest(x int, args ...string) {
	fmt.Println("SliceParamTest", x, args)

	for _, v := range args {
		fmt.Println("v=", v)
	}
}

// multiReturnValuesTest test
func multiReturnValuesTest() (int, string) {
	return 1, "multiReturnValuesTest"
}

// useMultiReturnValuesFunc used to test multiReturnValuesTest
func useMultiReturnValuesFunc(num int, str string) {
	fmt.Println(num, str)
}

// deferTest
func deferTest() {
	defer func() {
		fmt.Println("i am run in defer")
	}()
	var x int
	x = 1
	defer func() {
		x = 2
	}()
	fmt.Println("deferTest", x)
}

func panicTest() {
	panic("panicTest error")
}

func errorsTest(x int, y int) (int, error) {

	error := errors.New("division by zero")

	if y == 0 {
		return 0, error
	}
	z := x / y
	return z, nil
}
