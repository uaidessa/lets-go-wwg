package main

import "fmt"

func main() {

	var slice1 = []int{5, 6, 20, 40, 30, 1}
	var slice2 = []int{1, 27, 14, 230, 230, 89, 10}

	slice3 := append(slice1, slice2...)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
