package arraytest

import "fmt"

// TestArray test
func TestArray() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1", arr1)

	arr2 := [...]int{1: 10, 2: 3, 4: 50}
	fmt.Println("arr2", arr2)
	fmt.Println("arr2[1]", arr2[1])

	arr3 := [3]*int{1: new(int), 2: new(int)}
	fmt.Println("arr3", arr3)
	*arr3[1] = 1
	*arr3[2] = 100
	fmt.Println("arr3[1]", *arr3[1])
	fmt.Println("arr3[2]", *arr3[2])

	arr4 := [5]int{11, 22, 33, 44, 55}

	for index, v := range arr4 {
		fmt.Println("index", index, "value", v)
	}
}
