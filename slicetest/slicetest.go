package slicetest

import (
	"fmt"
)

// SliceCreate is for slice create
func SliceCreate() {

	slice1 := make([]int, 7)
	fmt.Println("slice1 is ", slice1)
	slice1[0] = 1
	fmt.Println("slice1 is ", slice1)
	fmt.Println("slice1[0] is ", slice1[0])
	fmt.Println("slice1 len is ", len(slice1))

	for index, val := range slice1 {
		fmt.Println("index ", index, "val ", val)
	}
}

// SliceAndVariadic for test slice and variadic
func SliceAndVariadic() {
	fmt.Printf("%.2f \n", average(10.0, 30, 40))
	slice1 := []float64{1, 3, 4, 5, 2}
	// slice + variadic
	fmt.Printf("%.2f \n", average(slice1...))
}

// average is for cal these num you give and return
func average(nums ...float64) float64 {

	var sum float64 = 0

	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))
}

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

	// test unInit slice
	var nilSlice []int
	fmt.Println("nil_slice", nilSlice)

	// test slice cut
	slice2 := make([]int, 6)
	// [0 0 0 0 0 0]
	fmt.Println("slice2", slice2)
	slice2[0] = 66
	slice2[5] = 55
	slice2[3] = 44
	// [66 0 0 44 0 55]
	fmt.Println("slice2", slice2)

	// [44 0 55]
	slice3 := slice2[3:]
	fmt.Println("slice3:", slice3)

	slice2[0] = 77
	slice3 = slice2[0:]
	// [77 0 0 44 0 55]
	fmt.Println("slice3", slice3)

	slice2[3] = 33
	// [77 0 0 33 0 55]
	fmt.Println("slice2", slice2)
	slice4 := slice2[3:4]
	// [33]
	fmt.Println("slice4", slice4)

	// [77 0 0 33 0 55]
	// [77 0 0]
	slice5 := slice2[:3]
	fmt.Println("slice5", slice5)
}
