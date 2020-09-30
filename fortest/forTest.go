package fortest

import "fmt"

func ForAppend() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func ForGetVar() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		//  v 是临时变量，所以newArr 是通一个值
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

// map 的遍历是无须的，因为遍历开始的时候，会随机 桶的位置
func ForMap() {
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		println(k, v)
	}
}
