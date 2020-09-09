package combiningtest

import "testing"

func TestAdmin(t *testing.T) {

	u := user{email: "554933654@qq.com", name: "Lucy"}

	a1 := admin{u, "100"}

	a2 := admin2{"100", u}

	a3 := new(admin2)
	a3.level = "1"
	a3.u = u

	SendNotify(u)
	SendNotify(a1)
	SendNotify(a2)
	SendNotify(a3)
}
