package queue

type Queue []int

// q 所指向的 Slice 改变
func (q *Queue) Push(v int) {
	*q = append(*q, v) // 不仅改里面的内容，它本身都能改掉
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
