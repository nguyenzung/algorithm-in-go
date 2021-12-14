package main

import "fmt"

func main() {
	i := 0
	for {
		if i == 20 {
			break
		} else {
			i++
		}
	}

	fmt.Printf(" i = %d", i)
}
