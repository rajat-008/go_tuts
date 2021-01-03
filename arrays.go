package main

import (
	f "fmt"
)



func main() {
	st := "ab"
	f.Println(st, len(st))
	st += "anmmf"
	f.Println(st, len(st))
}
