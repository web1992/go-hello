package recovertest

import "fmt"

// RecoverTest
func RecoverTest() {

	defer func() {
		fmt.Println("defer start >>>>>>>")
		if err := recover(); err != nil {
			fmt.Println("err is", err)
		}
		fmt.Println("defer end >>>>>>>>>")
	}()

	fmt.Println("RecoverTest run <<<<<<<<")
	panic("I am error")
	// below code not run because panic is rise
	fmt.Println("RecoverTest end <<<<<<<<")
}
