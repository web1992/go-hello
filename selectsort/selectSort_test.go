package selectsort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {

	arr := []int{1, 20, 30, 1, 45, 13, 00, 99}
	sortArr := selectSort(arr)
	fmt.Println("sorArr is ", sortArr)
}
func TestSelectSort2(t *testing.T) {

	arr := []int{1, 20, 30, 1, 45, 13, 00, 99}
	sortArr := selectSort2(arr)
	fmt.Println("sorArr is ", sortArr)

	arr2 := []int{41, 20, 30, 1, 45, 13, 00, 99, 89, -1, 123, 123999}
	sortArr2 := selectSort2(arr2)
	fmt.Println("sorArr2 is ", sortArr2)
}
