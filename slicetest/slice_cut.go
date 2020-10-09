package slicetest

import "fmt"

func SliceCut() {

	var intArr []int

	for i := 0; i < 10; i++ {
		intArr = append(intArr, i)
	}

	fmt.Println("intArr is", intArr)

	intArr2 := intArr[1:]
	fmt.Println("intArr2 is", intArr2)

}
