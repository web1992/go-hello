package mergesort

func merge(a []int, b []int) []int {

	t := make([]int, len(a)+len(b))

	ai := 0
	bi := 0

	for ai < len(a) && bi < len(b) {

		if a[ai] <= b[bi] {
			t[ai+bi] = a[ai]
			ai++
		} else {
			t[bi+ai] = b[bi]
			bi++
		}

	}

	for ai < len(a) {
		t[ai+bi] = a[ai]
		ai++
	}

	for bi < len(b) {
		t[bi+ai] = b[bi]
		bi++
	}
	return t
}

// 合并排序
func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	middle := len(arr) / 2

	a := MergeSort(arr[:middle])
	b := MergeSort(arr[middle:])

	return merge(a, b)
}
