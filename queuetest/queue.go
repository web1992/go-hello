package queuetest

import "container/list"

type Queue struct {
	l    *list.List
	size int
}

func NewQueue() *Queue {

	return &Queue{
		l:    list.New(),
		size: 0,
	}
}

func (q *Queue) Empty() bool {

	return q.size == 0
}

func (q *Queue) ArrEnter(arrE []interface{}) {

	for _, e := range arrE {
		q.Enter(e)
	}
}

func (q *Queue) Enter(e interface{}) {

	q.l.PushBack(e)
	q.size++
}

func (q *Queue) Out() (bool, interface{}) {

	if q.size == 0 {
		return false, nil
	}
	_e := q.l.Front()
	q.l.Remove(_e)
	q.size--
	return true, _e.Value
}
