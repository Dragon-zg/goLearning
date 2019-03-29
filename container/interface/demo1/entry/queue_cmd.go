package main

import (
	"goLearning/container/interface/demo1/queue"
	"fmt"
)

func main()  {
	queue := queue.Queue{1,2}
	queue.Push(500)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())

	fmt.Println("-----------------------------")
	queue.Push("abd")
	fmt.Println(queue.IsEmpty())
	queue.Push("@#$")
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())
}
