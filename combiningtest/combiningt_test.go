package combiningtest

import "testing"

func TestAdmin(t *testing.T) {

	u := user{email: "554933654@qq.com", name: "Lucy"}

	a := admin{u, "100"}
	a2 := admin2{"100", u}

	SendNotify(u)
	SendNotify(a)
	SendNotify(a2)
}
