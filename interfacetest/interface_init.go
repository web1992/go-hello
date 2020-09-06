package interfacetest

import "fmt"

type Phone interface {
	call(num string) error
}

type IOSPhone struct{}

type HuaWeiPhone struct{}

func (p HuaWeiPhone) call(num string) error {
	fmt.Println("HuaWeiPhone calling num is", num)
	return nil
}

func (p IOSPhone) call(num string) error {
	fmt.Println("IOSPhone calling num is", num)
	return nil
}

func Call() {
	var p IOSPhone
	//var p = new(IOSPhone)

	// var pMap = make(map[string]IOSPhone)
	//pMap["xx"] = p

	e := p.call("110")
	if e != nil {
		fmt.Println("call suc!")
	}
}
