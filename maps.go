package main

import "fmt"

func main() {
	mp := map[string]int{"hello": 0}
	mp["hello"] = 1
	v, f := mp["hello"]
	fmt.Println(v, f)
	v, f = mp["a"]
	fmt.Println(v, f)
}
