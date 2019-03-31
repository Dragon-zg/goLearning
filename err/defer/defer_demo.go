package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	//a()
	//func1("Go")
	func2()
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

func func2() {
	defer fmt.Println("this is first defer")
	defer fmt.Println("this is second defer")
}

func func3() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("this i is:", i)
		if i == 5 {
			panic("panic")
		}
	}

}
