package searchseller

import (
	"container/list"
	"fmt"
	"strings"
)

func testList() {

	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	len := l.Len()
	for i := 0; i < len; i++ {
		e := l.Front()
		l.Remove(e)
		fmt.Println(e.Value)
	}
}

/**
图的广度优先遍历
使用 map+queue 进行遍历
*/
func searchSeller() {

	graph := initMap()
	queue := list.New()

	for _, v := range graph["you"] {
		queue.PushBack(v)
	}
	searched := make(map[string]string)

	c := 0
	var searchArr []string

	for queue.Len() > 0 {

		c++
		e := queue.Front()
		queue.Remove(e)
		s := e.Value.(string)
		if _, ok := searched[s]; ok {
			fmt.Println("had searched", s)
			continue
		}
		searched[s] = s
		searchArr = append(searchArr, s)
		searchArr = append(searchArr, "---->")
		fmt.Println("check ", s)
		if isSeller(s) {
			fmt.Println("find is", s)
			fmt.Println("c is", c)
			fmt.Println(searchArr)
			return
		}
		arr := graph[s]
		for _, v := range arr {
			queue.PushBack(v)
		}
	}
	fmt.Println("c is", c)
	fmt.Println(searchArr)
}

func initMap() map[string][]string {
	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	return graph
}

func isSeller(s string) bool {
	return strings.HasSuffix(s, "m")
}
