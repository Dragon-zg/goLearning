package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(ii int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", ii)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
