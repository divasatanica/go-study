package main

import "fmt"

// Stack implemented by array
type Stack struct {
	data []int
}

// Push an item into the stack
func (s Stack) Push(data int) {
	stack := &s
	stackData := stack.data
	length := len(stackData)
	cap := cap(stackData)

	if outOfSize := length+1 >= cap; outOfSize {
		cap *= 2
		newData := make([]int, cap)
		for i, v := range stackData {
			newData[i] = v
		}
		s.data = newData
	}
	s.data[length] = data
}

func newStack() Stack {
	return Stack{make([]int, 1)}
}

func main() {
	stack := newStack()

	stack.Push(1)

	fmt.Println(stack.data)
}
