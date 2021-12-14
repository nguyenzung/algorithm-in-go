package main

import "fmt"

func testArray(arr *[]int) {
	*arr = (*arr)[1:]
}

func testMap(mapper *map[int]int) {
	(*mapper)[1] = 1
	m := make(map[int]int)
	mapper = &m
	(*mapper)[1] = 2

}

func main() {
	mapper := make(map[int]int)
	testMap(&mapper)
	fmt.Println("[Mapper]", mapper)

	arr := []int{1, 2, 3}
	testArray(&arr)
	fmt.Println("[Array]", arr)
}
