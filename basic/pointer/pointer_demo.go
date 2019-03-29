package main

import "fmt"

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	
}

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
	fmt.Printf("reply: %v \n", reply)
	tmp := 111
	reply = &tmp
	fmt.Printf("*reply: %v \n", *reply)
}