package main

import "fmt"

func main() {
	//demo2()
	//demo3()
	demo4()
}

func demo1() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
		//此处虽然重新赋值，但其实未生效
		season = "更改值."
	}
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}

//切片重组
func demo2() {
	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}
	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

//切片复制，追加
func demo3() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n ==

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

func demo4() {
	slice := []int{1, 2, 3, 4, 5}

	// 切除切片 a 中从索引 i 至 j 位置的元素： a = append(a[:i], a[j:]...)
	slice2 := append(slice[:1], slice[3:]...)
	fmt.Println(slice)
	fmt.Println(slice2)
	//此处修改了slice的值 取前3位给slice2,第一个元素的内存地址一样 就是 最好的证明
	fmt.Println(&slice[0])
	fmt.Println(&slice2[0])

	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))
}
