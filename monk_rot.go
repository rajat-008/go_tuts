/*
	// Sample code to perform I/O:
	//
	fmt.Scanf("%s", &myname)            // Reading input from STDIN
	fmt.Println("Hello", myname)        // Writing output to STDOUT
	//
	// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type myint int

func main() {
	var str string
	var i []int

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		str = reader.Text()
		flds := strings.Fields(str)
		for _, fld := range flds {
			k, _ := strconv.Atoi(fld)
			i = append(i, k)
		}
		fmt.Println(i)
	}
	fmt.Println("out")
}
