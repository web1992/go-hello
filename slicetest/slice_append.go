package slicetest

import "fmt"

func Append() {

	var s1 []int
	fmt.Println(s1)
	s1 = append(s1, 1)
	fmt.Println("after append", s1)

	s2 := make([]int, 7)
	s2[0] = 111
	s2[6] = 990
	fmt.Println("after s[0]", s2, " s2 len=", len(s2))

	s2 = append(s2, 666)

	fmt.Println("s2 is ", s2, " s2 len=", len(s2))

}
