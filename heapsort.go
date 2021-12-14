package main

import "fmt"

type Comparator func(interface{}, interface{}) bool

func heapsort(arr []interface{}, comp Comparator) {
	buildHeap(arr, comp)
	sort(arr, comp)
	revertArr(arr)
}

func buildHeap(arr []interface{}, comp Comparator) {
	for i := range arr {
		buildHeapWithIndex(arr, comp, i)
	}
}

// arr is already build with size - 1 in case size > 0
func buildHeapWithIndex(arr []interface{}, comp Comparator, index int) {
	parentIndex := getParent(index, arr)
	if parentIndex >= 0 {
		if !comp(arr[parentIndex], arr[index]) {
			swap(parentIndex, index, arr)
		}
		buildHeapWithIndex(arr, comp, parentIndex)
	}
}

func sort(arr []interface{}, comp Comparator) {
	for i := len(arr) - 1; i >= 0; i-- {
		sortWithIndex(arr, comp, i)
	}
}

func sortWithIndex(arr []interface{}, comp Comparator, index int) {
	if index == 0 {
		return
	} else {
		swap(0, index, arr)
		maintainHeap(arr, comp, index-1)
	}
}

func maintainHeap(arr []interface{}, comp Comparator, size int) {
	maintainHeapAtIndex(arr, comp, 0, size)
}

func maintainHeapAtIndex(arr []interface{}, comp Comparator, index int, size int) {
	leftChildIndex := getLeftChild(index, arr)
	if leftChildIndex != -1 && leftChildIndex <= size {
		lowest := leftChildIndex
		if comp(arr[index], arr[lowest]) {
			lowest = index
		}
		rightChildIndex := getRightChild(index, arr)
		if rightChildIndex != -1 && rightChildIndex <= size {
			if !comp(arr[lowest], arr[rightChildIndex]) {
				lowest = rightChildIndex
			}
		}
		if lowest != index {
			swap(index, lowest, arr)
			maintainHeapAtIndex(arr, comp, lowest, size)
		}
	}
}

func revertArr(arr []interface{}) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func swap(i int, j int, arr []interface{}) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func getParent(index int, arr []interface{}) int {
	return (index - 1) >> 1
}

func getLeftChild(index int, arr []interface{}) int {
	if (index << 1) + 1 < len(arr) {
		return (index << 1) + 1
	}
	return -1
}

func getRightChild(index int, arr []interface{}) int {
	if (index << 1) + 2 < len(arr) {
		return (index << 1) + 2
	}
	return -1
}

func main() {
	arr := []interface{}{7, 2, -8, 3, 4, -5, 0, 9, -1, 6, -10, -1, 5, -2}
	heapsort(arr, func(x interface{}, y interface{}) bool {
		return x.(int) < y.(int)
	})
	fmt.Println(arr)
}
