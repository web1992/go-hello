package dijkstra

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

// 以下实现优先级队列
// This example demonstrates a priority queue built using the heap interface.
// entry 是 priorityQueue 中的元素
// https://www.jianshu.com/p/15e6d436d471
type entry struct {
	key      string
	priority int
	// index 是 entry 在 heap 中的索引号
	// entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 entry.priority 一直不变的话，可以删除 index
	index int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) ToString() string {
	s := "["
	for _, v := range pq {
		s += fmt.Sprint(v.key, ",")
	}
	s = strings.TrimSuffix(s, ",")
	s += "]"
	return s
}

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PQ) PoPKey() (string, int) {
	e := pq.Pop()
	return e.(*entry).key, e.(*entry).priority
}

func (pq *PQ) PushKey(k string, p int) {
	pq.Push(&entry{key: k, priority: p})
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*entry)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
	sort.Sort(pq)

}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	sort.Sort(pq)
	return temp
}

// update modifies the priority and value of an entry in the queue.
func (pq *PQ) update(entry *entry, value string, priority int) {
	entry.key = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}