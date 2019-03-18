package main

import "fmt"

func main() {
	//demo2()
	//demo3()
	demo4()
}

func forStr() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d %d %U '%c' % X\n", index, rune, rune, rune, []byte(
			string(rune)))
	}
}

func demo1() {
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}
}

func demo2() {
	for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
	}
}
func demo3() {
	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
}
func demo4() {
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
