package functest

import "fmt"

// Main test
func Main() {

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

// mulitReturnValuesTest test
func mulitReturnValuesTest() (int, string) {
	return 1, "a"
}
