package functest

import "fmt"

// FuncTest test
// func as param
func FuncTest(msg string, f func(x, y int) int) {
	fmt.Println(msg, ":", f(1, 1))
}

// Add func
// return func
func Add() func(a, b int) int {

	f := func(a, b int) int {
		return a + b
	}
	return f
}
