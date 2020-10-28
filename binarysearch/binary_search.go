package binarysearch

import "fmt"

// binarySearch 二分查找
func binarySearch() {
	// low .........mid..............high
	// 数组必须有序
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)

	findNum := 1

	low := 0
	high := len(arr)
	c := 0
	for low <= high {
		c++
		mid := (high + low) / 2
		k := arr[mid]
		if findNum == k {
			fmt.Println("[find it, index is]:", mid)
			break
		} else if findNum > k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("[find times]:", c)

}
