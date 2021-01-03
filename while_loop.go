package main

import "fmt"

func main() {
	sum := 5
	for sum > 1 {
		fmt.Println(sum)
		sum -= 1
	}
}
