package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 5, 5)
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr = append(arr, arr2[:]...)
	fmt.Println(arr)

}
