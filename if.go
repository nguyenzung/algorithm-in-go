package main

import (
	"fmt"
)

func checkIf(x, n uint64) uint64 {
	if i := x; i < n {
		return i
	} else {
		return 0
	}

}

func main() {
	fmt.Println(
		checkIf(2, 3),
		checkIf(3, 3),
		checkIf(4, 3),
	)
}
