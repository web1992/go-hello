package quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 50, 40, 30, 20, 10}
	arr = quickSort(arr)
	fmt.Println(arr)
}
