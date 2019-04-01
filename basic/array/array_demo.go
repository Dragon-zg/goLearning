package main

import (
	"fmt"
)

func main() {
	//init1()
	//fmt.Println("-----------")
	//init2()
	//fmt.Println("-----------")
	//init3()
	//fmt.Println("-----------")
	fmt.Println(twoSum([] int{3,2,4}, 6))
}

func init1() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	for i := range arrAge {
		fmt.Println("Array item ", i, "is", arrAge[i])
	}
}

func init2() {
	arrLazy := [...]string{"a", "b", "c", "d"}
	for i := range arrLazy {
		fmt.Println("Array item", i, "is", arrLazy[i])
	}
}

func init3() {
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	for i := range arrKeyValue {
		fmt.Println("Array item", i, "is", arrKeyValue[i])
	}
}

func twoSum(nums []int, target int) []int {
	for idx1, v1 := range nums {
		for idx2, v2 := range nums {
			if idx1!= idx2 && v1 + v2 == target {
				var idx = make([]int, 2)
				idx[0] = idx1
				idx[1] = idx2
				return idx
			}
		}
	}

	panic("No two sum solution")
}