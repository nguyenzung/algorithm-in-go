package main

import "fmt"

func testQueue() {
	queue := make([]int, 0)
	for i := 0; i < 10; i++ {
		queue = append(queue, i)
	}
	fmt.Println("Queue:", queue)
	queue = queue[1:]
	fmt.Println("Queue:", queue)
}

func testStack() {
	stack := make([]int, 0)
	for i := 0; i < 10; i++ {
		stack = append(stack, i)
	}
	fmt.Println("Stack:", stack)
	stack = stack[0 : len(stack)-1]
	fmt.Println("Stack:", stack)
}

func main() {
	testQueue()
	testStack()
}
