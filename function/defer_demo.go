package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	//a()
	func1("Go")
}

//defer 相当于java中finally
func a() {
	i := 0
	defer fmt.Println("defer:", i)
	i++
	fmt.Println("after defer:", i)
	return
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
