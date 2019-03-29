package queue

//切片起别名
type Queue []interface{}

//添加一个元素
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

//删除并返回一个元素
func (q *Queue) Pop() interface{} {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

//判断是否是空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
