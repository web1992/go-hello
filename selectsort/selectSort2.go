package selectsort

// selectSort2 选择排序
func selectSort2(arr []int) []int {

	l := len(arr)

	for i := 0; i < l; i++ {
		minIndex := findMin(arr, i)
		swap(arr, minIndex, i)
	}

	return arr
}

// swap 交换位置
// 把小的发放到初始位置上
func swap(arr []int, k, j int) {

	t := arr[k]
	arr[k] = arr[j]
	arr[j] = t

}

// findMin 找到一个最小的值索引
func findMin(arr []int, start int) int {

	minIndex := start
	minNum := arr[minIndex]
	for s := minIndex; s < len(arr)-1; s++ {

		if minNum > arr[s+1] {
			minNum = arr[s+1]
			minIndex = s + 1
		}
	}
	return minIndex
}
