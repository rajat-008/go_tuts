package main

import "fmt"

type mystruct struct {
	a int
}

func main() {
	var a_my mystruct
	a_my.a = 0
	fmt.Println(a_my.a)
	(&a_my).a = 1
	fmt.Println(*(&a_my))
	fmt.Println(&a_my == nil)
	a := " "
	fmt.Println(a == "")
}
