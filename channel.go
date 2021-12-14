package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	fmt.Println("Start sum")
	sum := 0
	for _, v := range s {
		time.Sleep(100 * time.Millisecond)
		sum += v
	}
	c <- sum // send sum to c
	fmt.Println("End sum")
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 2)
	sum(s[:len(s)/2], c)
	sum(s[len(s)/2:], c)

	fmt.Println("Push data to variable via channel")
	fmt.Printf("Pushed data to variable via channel: %d \n", <-c)
	fmt.Printf("Pushed data to variable via channel: %d \n", <-c)
	// fmt.Printf("Pushed data to variable via channel: %d \n", <-c)
}
