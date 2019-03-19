package main

import "fmt"

func main() {
	a()
}

//defer 相当于java中finally
func a() {
	i := 0
	defer fmt.Println("defer:", i)
	i++
	fmt.Println("after defer:", i)
	return
}
