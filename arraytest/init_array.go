package arraytest

import "fmt"

type User struct {
}

func InitArray() {

	var a [3]int
	fmt.Println("a", a)
	fmt.Println("&a", &a)

	var user User

	fmt.Println("user", user)
	fmt.Println("&user", &user)

	var a3 = [...]int{1, 3, 4}

	fmt.Println("a3", a3, "le is ", len(a3))

}
