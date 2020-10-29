package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// 数组必须有序
	arr := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}

	var rsIndex = binarySearch(4, arr)
	fmt.Println("rsIndex is ", rsIndex)

	rsIndex = binarySearch(0, arr)
	fmt.Println("rsIndex is ", rsIndex)
}
