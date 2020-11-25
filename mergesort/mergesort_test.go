package mergesort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 10, 9, 8, 7, 6}
	sortArr := MergeSort(arr)
	fmt.Println(sortArr)
}
