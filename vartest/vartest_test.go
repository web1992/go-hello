package vartest

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	VarTest()
}

func TestUser(t *testing.T) {

	var u User
	fmt.Println(u)
	fmt.Println(*(&u))
	fmt.Println("name is", u.Name)

	u2 := new(User)
	fmt.Println(&u2)
	fmt.Println("name is", u2.Name)

	u3 := &User{}
	fmt.Println(&u3)

}
