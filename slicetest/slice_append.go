package slicetest

import "fmt"

func Append() {

	var s1 []int
	fmt.Println(s1)
	s1 = append(s1, 1)
	fmt.Println("after append", s1, "ls is ", len(s1), "cap is ", cap(s1))

	s2 := make([]int, 7, 8)
	s2[0] = 111
	s2[6] = 990
	fmt.Println("after s[0]", s2, " s2 len=", len(s2), "cap is ", cap(s2))

	s2 = append(s2, 666)

	fmt.Println("s2 is ", s2, " s2 len=", len(s2), "cap is", cap(s2))
	// OutPut:
	// s2 is  [111 0 0 0 0 0 990 666]  s2 len= 8 cap is 8

	s2 = append(s2, 999)

	fmt.Println("s2 is ", s2, " s2 len=", len(s2), "cap is", cap(s2))
	// OutPut:
	// s2 is  [111 0 0 0 0 0 990 666 999]  s2 len= 9 cap is 16

}
