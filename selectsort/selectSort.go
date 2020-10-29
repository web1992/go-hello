package selectsort

// selectSort 选择排序
func selectSort(arr []int) []int {

	arrS := arr
	l := len(arrS)
	var rs []int
	for i := 0; i < l; i++ {
		si := findSmallest(arrS)
		rs = append(rs, arrS[si])
		arrS = pop(arrS, si)
	}

	return rs
}

// pop 删除指定索引处的元素，并返回新的slice
func pop(arr []int, index int) []int {
	var a []int

	for i, v := range arr {
		if i == index {
			continue
		}
		a = append(a, v)
	}

	return a
}

// findSmallest 找到 slice  中最新的元素
func findSmallest(arr []int) int {

	si := 0
	s := arr[0]
	for i, v := range arr {
		if s > v {
			s = v
			si = i
		}
	}
	return si
}
