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

	// test append
	slice1 = append(slice1, 66)
	fmt.Println("slice1 after append", slice1)

	// this will make error
	//slice1[4] = 66
	//fmt.Println("slice1", slice1)

	// test uninit slice
	var nilSlice []int
	fmt.Println("nil_slice", nilSlice)

	// test slice cut
	slice2 := make([]int, 6)
	// [0 0 0 0 0 0]
	fmt.Println("slice2", slice2)
	slice2[0] = 66
	// []
	fmt.Println("slice2", nilSlice)

	// [0 0 0]
	slice3 := slice2[3:]
	fmt.Println("slice3", slice3)

	slice2[0] = 77
	slice3 = slice2[0:]
	// [77 0 0 0 0 0]
	fmt.Println("slice3", slice3)

	slice2[3] = 33
	// [77 0 0 33 0 0]
	fmt.Println("slice2", slice2)
	slice4 := slice2[3:4]
	// [33]
	fmt.Println("slice4", slice4)

}
