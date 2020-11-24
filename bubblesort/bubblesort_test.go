package bubblesort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := bubbleSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	fmt.Println(arr)
}
