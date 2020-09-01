package queue

// type Queue []int
// 让这个Slice支持任何类型
type Queue []interface{}

// q 所指向的 Slice 改变
// 限制只能 push int
func (q *Queue) Push(v int) {
	*q = append(*q, v) // 不仅改里面的内容，它本身都能改掉
}

// func (q *Queue) Push(v interface{}) {
// 	*q = append(*q, v.(int)) //在里面限定，代价是运行时错误
// }

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int) // 这里要强制转换
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
