package arraytest

import "fmt"

// TestArray test
func TestArray() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1", arr1)

	arr2 := [...]int{1: 10, 2: 3, 4: 50}
	fmt.Println("arr2", arr2)
	fmt.Println("arr2[1]", arr2[1])

	arr3 := [3]*int{1: new(int), 2: new(int), 0: new(int)}
	fmt.Println("arr3", arr3, "arr3 len is", len(arr3))
	fmt.Println("*arr3[1] is", *arr3[1])
	fmt.Println("*arr3[2] is", *arr3[2])
	fmt.Println("*arr3[0] is", *arr3[0])
	*arr3[1] = 1
	*arr3[2] = 100
	fmt.Println("*arr3[1]", *arr3[1])
	fmt.Println("*arr3[2]", *arr3[2])

	arr4 := [5]int{11, 22, 33, 44, 55}

	for index, v := range arr4 {
		fmt.Print("index=", index, ",", "value=", v, "\t")
	}

	fmt.Println()

	arr5 := [3]string{}
	arr5[0] = "a"
	arr5[1] = "b"
	arr5[2] = "c"
	//arr5[3]="c"
	fmt.Println("arr5 is", arr5)
}
