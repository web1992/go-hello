package sorttest

import (
	"fmt"
	"sort"
)

// TestSort TestSort
// sort test
func TestSort() {

	arr := []int{10, 33, 20, 40, 66}

	sortfunc := func(a, b int) bool {
		return arr[a] < arr[b]
	}
	sort.SliceStable(arr, sortfunc)
	fmt.Println(arr)

}
