package main

import "fmt"

// Queue - Implemented using array
type Queue struct {
	data   []int
	length int
}

// Push - Push an item into the queue
func (q *Queue) Push(data int) {
	q.data = append(q.data, data)
	q.length++
}

// Shift - Remove the head element of the queue
func (q *Queue) Shift() {
	q.data = q.data[:len(q.data)-1]
	q.length--
}

// NewQueue - Get a new queue
func NewQueue() Queue {
	var empty []int
	return Queue{empty, len(empty)}
}

func main() {
	queue := NewQueue()
	q := &queue

	q.Push(3)
	q.Push(2)
	fmt.Println(q, q.length)
	q.Shift()
	fmt.Println(q, q.length)
}
