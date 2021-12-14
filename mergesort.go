package main

import "fmt"

type Comparator func(interface{}, interface{}) bool

func mergesort(arr []interface{}, comp Comparator) {
	sortWithIndex(arr, 0, len(arr), comp)
}

func sortWithIndex(arr []interface{}, start int, end int, comp Comparator) {
	if start+1 < end {
		middle := (start + end) >> 1
		sortWithIndex(arr, start, middle, comp)
		sortWithIndex(arr, middle, end, comp)
		merge(arr, start, middle, middle, end, comp)
	}
}

func merge(arr []interface{}, start1 int, end1 int, start2 int, end2 int, comp Comparator) {
	size := (end2 - start2) + (end1 - start1)
	tempArr := make([]interface{}, size)
	idx1, idx2, tempIdx := start1, start2, 0
	for idx1 < end1 && idx2 < end2 {
		if comp(arr[idx1], arr[idx2]) {
			tempArr[tempIdx] = arr[idx1]
			idx1++
		} else {
			tempArr[tempIdx] = arr[idx2]
			idx2++
		}
		tempIdx++
	}
	if idx1 < end1 {
		copy(tempArr[tempIdx:size], arr[idx1:end1])
	} else if idx2 < end2 {
		copy(tempArr[tempIdx:size], arr[idx2:end2])
	}
	copy(arr[start1:end1], tempArr[0:(end1-start1)])
	copy(arr[start2:end2], tempArr[(end1-start1):size])
}

func makeComp(comp func(int, int) bool) func(int, int) bool {
	comparator := func(x, y int) bool {
		return x < y
	}
	if comp != nil {
		comparator = comp
	}
	return comparator
}

func main() {
	fmt.Println("Merge Sort")
	arr := []interface{}{7, 2, -8, 3, -4, 4, -5, 0, 9, -1, 6, -10, -1, 5, -2}

	mergesort(arr, func(x interface{}, y interface{}) bool {
		return x.(int) < y.(int)
	})
	fmt.Println(arr)
	mergesort(arr, func(x interface{}, y interface{}) bool {
		return x.(int) > y.(int)
	})
	fmt.Println(arr)
}
