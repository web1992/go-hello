package fizzbuzztest

import "fmt"

/**
编写一个将数字从 1 打印到 100 的程序。
对于 3 的倍数就打印 ‘Fizz’ ，对于 5 的倍数 就打印 ‘Buzz’。
对于既是 3 又是 5 的倍数，就打印 ‘FizzBuzz’。
*/
func main() {

	i := 1
	for {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, "FizzBuzz")
			}
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			// nop
		}

		if i == 100 {
			break
		}
		i++
	}
}
