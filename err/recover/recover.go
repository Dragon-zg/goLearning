package main

import (
	"fmt"
)

func main() {
	tryRecover()
}

func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("nothing recover")
			return
		}
		//panic(r)
		if err, ok := r.(error); ok {
			fmt.Println("Error is:", err.Error())
		} else {
			panic(fmt.Sprintf("i don't konw what to do: %v", r))
		}
	}()
	//panic(errors.New("this is an error"))

	//a, b := 5, 0
	//fmt.Println(a/b)

	panic(123)
}
