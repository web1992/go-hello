package combiningtest

import "fmt"

type Task struct {
	id   int32
	name string
	no   string
	TaskFeature
}

type TaskFeature struct {
	num   int32
	price int32
	desc  string
}

func main() {
	task := Task{}

	task.name = "PLP"
	task.id = 1
	task.no = "OX000001"
	task.num = 10
	task.price = 200
	task.desc = "demo"

	fmt.Println(task)
}
