package main

import "fmt"

func hello12(c chan int) {
	i := 0

	for i < 10 {
		c <- i
		i += 1
	}
	return
}

func main() {
	ch := make(chan int)
	defer close(ch)
	go hello(123ch)
	for i := range ch {
		fmt.Println(i)
	}

}
