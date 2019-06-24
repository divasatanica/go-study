package main

import "fmt"

// Stack implemented by array
type Stack struct {
	data []int
}

// Push an item into the stack
func (s Stack) Push(pt *Stack, data int) {
	pt.data = append(s.data, data)
}

// Pop an item from the stack
func (s Stack) Pop(pt *Stack) {
	pt.data = pt.data[0:len(pt.data)-1]
}

func newStack() Stack {
	var empty []int
	return Stack{empty}
}

func main() {
	stack := newStack()
	pt := &stack
	stack.Push(pt, 1)
	stack.Push(pt, 2)
	fmt.Println(stack)
	stack.Pop(pt)
	fmt.Println(stack)
}
