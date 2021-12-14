package main

import "fmt"

func main() {
	a := make([]uint64, 3)
	a[0] = 11
	a[1] = 22
	a[2] = 33
	fmt.Printf(" Size of a: %d \n", len(a))

	for i, value := range a {
		fmt.Printf(" Range: %d %d \n", i, value)
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf(" a[%d] = %d", i, a[i])
	}
}
