package dijkstra

type Set struct {
	m map[string]string
}

type Node struct {
	name  string
	score int
}

func NewSet() *Set {
	return &Set{
		m: make(map[string]string),
	}
}

func (s *Set) Has(e string) bool {
	_, ok := s.m[e]
	return ok
}
func (s *Set) Add(e string) {
	s.m[e] = ""
}
