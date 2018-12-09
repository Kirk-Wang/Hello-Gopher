package queue

// An FIFO queue.
type Queue []interface{}

// Pushes the element into the queue.
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pops elements from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
