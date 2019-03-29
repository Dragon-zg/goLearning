package main

import (
	"fmt"
	"time"
)

//斐波那契数列 1 1 2 3 5 8 13 21 34 55 89
func main() {
	start := time.Now()
	result := 0
	for i := 0; i <= 41; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
