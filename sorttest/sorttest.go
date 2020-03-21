package sorttest

import (
	"fmt"
	"sort"
)

// TestSort TestSort
// sort test
func TestSort() {

	arr := []int{10, 33, 20, 40, 66}

	sortFunc := func(a, b int) bool {
		return arr[a] < arr[b]
	}
	sort.SliceStable(arr, sortFunc)
	fmt.Println(arr)

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)

}
