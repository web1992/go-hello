package stacktest

import "container/list"

type Stack struct {
	l    *list.List
	size int
}

func New() *Stack {

	return &Stack{
		l:    list.New(),
		size: 0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Push(e interface{}) {
	s.l.PushFront(e)
	s.size++
}

func (s *Stack) Pop() (bool, interface{}) {
	if s.size == 0 {
		return false, nil
	}
	_e := s.l.Front()
	s.l.Remove(_e)
	s.size--
	return true, _e.Value
}
