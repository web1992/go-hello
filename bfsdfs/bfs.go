package bfsdfs

import (
	"fmt"
	"web1992/queuetest"
)

// bfsdfs.png
// B,C 是A 的一度关系
// D,E 是A 的二度关系
// F   是A 的三读关系
func BFS() {

	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "C"},
		"C": {"B", "D", "E"},
		"D": {"B", "C", "E", "F"},
		"E": {"C", "D"},
		"F": {"D"},
	}
	fmt.Printf("%s\n", graph)
	queue := queuetest.NewQueue()
	checkedMap := map[string]string{}

	arr := graph["A"]
	arrEnter(queue, arr)

	for !queue.Empty() {

		_, e := queue.Out()
		eStr := e.(string)
		if _, ok := checkedMap[eStr]; ok {
			continue
		}
		fmt.Println("element is", eStr)
		checkedMap[eStr] = eStr
		arr2 := graph[eStr]
		arrEnter(queue, arr2)
	}

}

// Expect Output:
// map[A:[B C] B:[D C] C:[B D E] D:[C E F] E:[C D] F:[D]]
// element is B
// element is C
// element is D
// element is E
// element is F

func arrEnter(q *queuetest.Queue, arr []string) {
	for _, v := range arr {
		q.Enter(v)
	}
}
