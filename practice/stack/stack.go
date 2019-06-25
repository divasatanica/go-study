package main

import "fmt"

// Stack implemented by array
type Stack struct {
	data []int
}

// Push an item into the stack
func (s *Stack) Push(data int) {
	s.data = append(s.data, data)
}

// Pop an item from the stack
func (s *Stack) Pop() {
	s.data = s.data[0 : len(s.data)-1]
}

func newStack() Stack {
	var empty []int
	return Stack{empty}
}

func main() {
	stack := newStack()
	pt := &stack
	pt.Push(1)
	pt.Push(2)
	fmt.Println(stack)
	pt.Pop()
	fmt.Println(stack)
}
