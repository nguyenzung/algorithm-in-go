package main

import (
	"fmt"
	"runtime"
)

func swap(x, y *uint64) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	var i, j uint64
	i, j = 12, 21
	var p *uint64
	p = &i // point to i
	*p++

	fmt.Printf("[i = %d | *p = %d | p = %p] \n", i, *p, p)

	fmt.Printf("Before Swap [i = %d | j = %d ] \n", i, j)

	swap(&i, &j)
	fmt.Printf("After Swap [i = %d | j = %d ] \n", i, j)

	fmt.Printf("Runtime %s \n\n", runtime.GOARCH)

	var a [10]uint64
	for index := 0; index < len(a); index++ {
		fmt.Printf("Address of index: %d is %p \n", index, &a[index])
	}

}
