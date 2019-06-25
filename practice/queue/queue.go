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

	queue.Push(3)
	queue.Push(2)
	fmt.Println(queue, queue.length)
	queue.Shift()
	fmt.Println(queue, queue.length)
}
