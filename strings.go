package main

import "fmt"

func main() {
	str := "a"
	var b string
	fmt.Scanf("%s", &b)
	str += "dfff"
	fmt.Println(str[:len(str)-1] + b)
	mt := ""
	fmt.Println(len(mt))
	mt += "htmml"
	fmt.Println(len(mt))
}
