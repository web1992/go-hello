package slicetest

import (
	"fmt"
)

// SliceTest test
func SliceTest() {

	slice1 := make([]int, 3, 5)

	fmt.Println("slice1", slice1)
	fmt.Println("slice1 len", len(slice1))
	fmt.Println("slice1 cap", cap(slice1))
	slice1[0] = 1
	fmt.Println("slice1", slice1)

	slice1 = append(slice1, 66)
	fmt.Println("slice1", slice1)

	// this will make error
	//slice1[4] = 66
	//fmt.Println("slice1", slice1)

	// test uninit slice
	var nilSlice []int
	fmt.Println("nil_slice", nilSlice)

	// test slice cut
	slice2 := make([]int, 6)
	fmt.Println(slice2)
	slice2[0] = 66

	// [0 0 0]
	slice3 := slice2[3:]
	fmt.Println(slice3)

	// [77 0 0 0 0 0]
	slice2[0] = 77
	slice3 = slice2[0:]
	fmt.Println(slice3)

	// [33]
	slice2[3] = 33
	slice4 := slice2[3:4]
	fmt.Println(slice4)

}
