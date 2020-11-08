package bfsdfs

import (
	"fmt"
	"web1992/stacktest"
)

// 深度优先遍历
func DFS() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "C"},
		"C": {"B", "D", "E"},
		"D": {"C", "E", "F"},
		"E": {"C", "D"},
		"F": {"D"},
	}
	fmt.Printf("%s\n", graph)

	stack := stacktest.New()
	stackPush(stack, graph["A"])

	checkedMap := make(map[string]string)

	for !stack.Empty() {
		_, e := stack.Pop()
		es := e.(string)

		if _, ok := checkedMap[es]; ok {
			continue
		}
		checkedMap[es] = es
		fmt.Println("element is ", es)
		stackPush(stack, graph[es])
	}

}

// Expect Output:
// map[A:[B C] B:[D C] C:[B D E] D:[C E F] E:[C D] F:[D]]
// element is  C
// element is  E
// element is  D
// element is  F
// element is  B

func stackPush(stack *stacktest.Stack, arr []string) {

	for _, v := range arr {
		stack.Push(v)
	}
}
