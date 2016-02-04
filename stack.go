package gotastructures

type Stack struct {
	top *StackNode
}

type StackNode struct {
	data interface{}
	next *StackNode
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.top.data
}

func (s *Stack) Push(data interface{}) {
	s.top = &StackNode{data, s.top}
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	data := s.top.data
	s.top = s.top.next
	return data
}
