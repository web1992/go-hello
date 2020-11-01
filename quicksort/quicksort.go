package quicksort

// 快速排序
func quickSort(arr []int) []int {
	// 结束条件
	if len(arr) <= 1 {
		return arr
	}
	// 基准值
	pivot := arr[0]
	var less []int
	var greater []int

	for _, v := range arr[1:] {
		if v > pivot {
			greater = append(greater, v)
		} else {
			less = append(less, v)
		}
	}

	// 递归排序
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}
