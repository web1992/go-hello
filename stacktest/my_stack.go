package stacktest

type MyStack struct {
	elements []string
	size     int
	pointer  int
}

func NewStack() *MyStack {

	_size := 100
	return &MyStack{
		elements: make([]string, _size),
		size:     _size,
		pointer:  0,
	}
}

func (s *MyStack) log() {
	for i := 0; i < s.pointer; i++ {
		println(s.elements[i])
	}

}
func (s *MyStack) pop() string {

	if s.pointer == 0 {
		return "nil"
	}
	//println("pop s.pointer", s.pointer)
	s.pointer--
	t := s.elements[s.pointer]
	return t
}

func (s *MyStack) push(element string) {

	//println("push ", s.pointer)
	if s.pointer == s.size {
		panic("stack is full")
	}
	s.elements[s.pointer] = element
	//println(s.pointer)

	s.pointer++

}
