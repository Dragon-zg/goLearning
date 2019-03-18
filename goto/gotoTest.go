package main

import "fmt"

func main() {
	//label1()
	//label2()
	//gotoDemo()
	//demo1()
	demo2()
}

func label1() {
	//标签的作用对象为外部循环
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
func label2() {
	//标签的作用对象为外部循环
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

func gotoDemo() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

func demo1()  {
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 { break }
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++;
	}
	fmt.Println("A statement just after for loop.")
}

func demo2(){
	for i := 0; i<7 ; i++ {
		if i%2 == 0 { continue }
		fmt.Println("Odd:", i)
	}
}