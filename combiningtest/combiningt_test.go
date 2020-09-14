package combiningtest

import (
	"fmt"
	"testing"
)

func TestAdmin(t *testing.T) {

	u := user{email: "554933654@qq.com", name: "Lucy"}

	a1 := admin{u, "100"}

	a2 := admin2{"100", u}

	a3 := new(admin2)
	a3.level = "1"
	a3.u = u
	fmt.Printf("%s", u.name)
	fmt.Printf("%s", u.email)

	SendNotify(u)
	SendNotify(a1)
	SendNotify(a2)
	SendNotify(a3)
}

// test use var init
func TestAdmin3(t *testing.T) {
	var admin3 admin2
	admin3.notify()

	admin4 := new(admin2)
	admin4.u = user{"ManFen", "110@qq.com"}
	admin4.level = "666"
	admin4.notify()
}
