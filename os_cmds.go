package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("%v", os.Args[0:])
	a, _ := exec.LookPath("python")
	fmt.Println("%v", a)
}
