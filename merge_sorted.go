package main

import "fmt"

func main() {
	A := []int{1, 4, 5, 6, 2, 4, 6}
	pt := 4
	ct := 0
	for i, num := range A[:len(A)-1] {
		for j, num2 := range A[pt:] {
			ct += 1
			if num <= num2 && i < pt {
				continue

			}
			fmt.Println(A)
			A[i], A[pt+j] = A[pt+j], A[i]
			fmt.Println(A)
			break
		}
	}
	fmt.Println(A)
	fmt.Println(ct)

}
