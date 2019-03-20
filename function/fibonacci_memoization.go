package main

import (
	"time"
	"fmt"
)

const LIM = 41

var fibs [LIM]uint64

//斐波那契数列 代码优化 1 1 2 3 5 8 13 21 34 55 89
func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci_other(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci_other(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci_other(n-1) + fibonacci_other(n-2)
	}
	fibs[n] = res
	return
}