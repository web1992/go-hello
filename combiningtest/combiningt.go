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
	user  user
}

func (u user) notify() {
	fmt.Println("Notify by user name=", u.name, " email=", u.email)
}

func (a admin2) notify() {
	fmt.Println("Notify by admin2 name=", a.user.name, " email=", a.user.email)
}

func SendNotify(n notifier) {
	n.notify()
}
