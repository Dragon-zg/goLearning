package main

import "fmt"

func main() {
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
