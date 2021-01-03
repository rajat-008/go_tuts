3
5 2
1 2 3 4 5
4 2
2 4 7 6
10 11
1 2 3 4 5 6 7 8 9 12
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
	"strings"
)

func main() {
	var t int
	var n, k int
	var str string
	res := ""
	fmt.Scanf("%d", &t)
	reader := bufio.NewScanner(os.Stdin)
	i := 0
	for i < t {
		fmt.Scanf("%d %d\n", &n, &k)
		reader.Scan()
		str = reader.Text()
		fields := strings.Fields(str)
		// fmt.Println(fields)

		nr := k % n
		//fmt.Println("nr=", nr)

		for _, ele := range fields[len(fields)-nr:] {
			res += (ele + " ")

		}
		// fmt.Println(fields[nr:])
		// fmt.Println(fields[0:nr])

		for _, ele := range fields[0 : len(fields)-nr] {
			res += (ele + " ")

		}

		res += "\n"
		i += 1
		// fmt.Println(fields)

	}
	fmt.Println(res)

}
