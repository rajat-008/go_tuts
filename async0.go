//concurrency

package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println("hello")
	time.Sleep(5000)
	go fmt.Println("1234")
	time.Sleep(5000)
}
