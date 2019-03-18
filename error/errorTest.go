package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	orig := "ABC"
	//origInt := "64"
	// var an int
	var newS string
	// var err error
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	//an, err := strconv.Atoi(origInt)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		fmt.Println("err:", err)
		os.Exit(1)
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
