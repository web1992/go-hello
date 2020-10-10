package stringtest

import "fmt"

// StringTest test
func StringTest() {

	// change string
	s := "www.web1992.cn"
	fmt.Println(s)
	arr := []byte(s)
	arr[0] = 'b'
	s = string(arr)
	fmt.Println(s)

}

func StringToRune() {
	str := "Hello 你好"
	r := []rune(str) // 8
	fmt.Println("len is", len(r))
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}

	fmt.Println()
	for _, v := range r {
		fmt.Printf("%c", v)
	}
}
