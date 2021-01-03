package main

import "fmt"

func main() {
	arr := make([]int, 10, 10)
	arr2 := arr[:]
	fmt.Println(arr, len(arr), cap(arr))
	var i int
	i = 0
	for i < 5 {
		arr[i] = 1
		i += 1
	}
	arr = arr[:0]
	fmt.Println(arr, len(arr), cap(arr))
	i = 0
	for i < 5 {
		arr = append(arr, 10-i)
		i += 1
	}
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(arr, len(arr2), cap(arr2))
}
