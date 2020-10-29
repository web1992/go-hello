package binarysearch

import "fmt"

// binarySearch 二分查找
func binarySearch(findNum int, arr []int) int {
	// low .........mid..............high
	// arr := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}

	low := 0
	high := len(arr)
	c := 0
	for low <= high {
		c++
		mid := (high + low) / 2
		k := arr[mid]
		if findNum == k {
			fmt.Println("[find it, index is]:", mid)
			return mid
		} else if findNum > k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("[find times]:", c)

	return -1
}
