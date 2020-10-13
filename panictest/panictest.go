package panictest

import (
	"fmt"
)

func panicError() {
	panic("!!! I am panic !!!")
}

func main() {
	fmt.Println("main start")
	panicRun()
	// 使用了 recover 捕获了 panic
	// 所以下面的代码依然会执行
	fmt.Println("main end")
}

func panicRun() {

	defer func() {
		// recover 必须在 defer 中使用
		if err := recover(); err != nil {
			fmt.Println("panic recover")
			fmt.Println(err.(string)) // 将 interface{} 转型为具体类型。
		}

	}()

	panicError()
	fmt.Println("this will not run")
}
func main2() {
	fmt.Println("main2 start")
	panicRun2()
	fmt.Println("main2 end")
}

func panicRun2() {
	panicError()
	fmt.Println("this will not run")
}
