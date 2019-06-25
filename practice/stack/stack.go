package main

import "fmt"

// Stack implemented by array
type Stack struct {
	data   []int
	length int
}

// Push an item into the stack
func (s *Stack) Push(data int) {
	s.data = append(s.data, data)
	s.length++
}

// Pop an item from the stack
func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
	s.length--
}

// GetTop - Get stack's top element
func (s *Stack) GetTop() int {
	length := s.length
	return s.data[length-1]
}

// NewStack - Get a new stack
func NewStack() Stack {
	var empty []int
	return Stack{empty, len(empty)}
}

func main() {
	stack := NewStack()
	stack.Push(5)
	stack.Push(2)
	fmt.Println(stack, stack.length)
	fmt.Println(stack.GetTop())
	stack.Pop()
	fmt.Println(stack, stack.length)
	fmt.Println(stack.GetTop())
}
