package stringtest

import "fmt"

// StringTest test
func StringTest() {
	str := "Hello 你好"
	r := []rune(str) // 8
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}

	// change string
	s := "www.web1992.cn"
	fmt.Println(s)
	arr := []byte(s)
	arr[0] = 'b'
	s = string(arr)
	fmt.Println(s)

}
