package main

import (
	"fmt"
	"os"
)

func main() {
	for _, varr := range os.Args {
		fmt.Printf("%s\n", varr)
	}
}
