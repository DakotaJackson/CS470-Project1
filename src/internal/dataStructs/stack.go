package dataStructs

// Following is a basic stack with basic functionality, following
// 	the LIFO standard.

// Stack is the data structure used for different pathing algorithms.
type Stack struct {
	length int
	top    *nodeS
}

// node is used in the stack data structure (only using ints for the map)
type nodeS struct {
	value int
	prev  *nodeS
}

// InitStack creates a blank stack for use elsewhere.
func InitStack() *Stack {
	return &Stack{
		0,
		nil,
	}
}

// Push adds a node to the top of the stack.
func (s *Stack) Push(value int) {
	n := &nodeS{value, s.top}
	s.top = n
	s.length++
}

// Pop removes a node from the top of the stack.
// 	It also returns false and 0 if the stack has no nodes.
func (s *Stack) Pop() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value, true
}

// GenLenS returns the length of the stack.
func (s *Stack) GenLenS() int {
	return s.length
}
