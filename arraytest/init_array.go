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

}
