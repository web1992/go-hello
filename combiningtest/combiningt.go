package combiningtest

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

type admin2 struct {
	level string
	u     user
}

func (u user) notify() {
	fmt.Printf("Notify by user name=%s email=%s \n", u.name, u.email)
}

func (a admin2) notify() {
	fmt.Printf("Notify by admin name=%s email=%s level=%s \n", a.u.name, a.u.email, a.level)
}

func SendNotify(n notifier) {
	n.notify()
}
