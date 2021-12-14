package main

import (
	"fmt"
	"go-avltree"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	str := strconv.Itoa(int(ip[0]))
	if len(ip) == 4 {
		for i := 1; i < 4; i++ {
			str += ("." + strconv.Itoa(int(ip[i])))
		}
	}
	return str
}

func main() {

	var i int = 0
	i = (i - 1) >> 1
	fmt.Println(i)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	tree := new(avltree.AVLTree)
	keys := []int{3, 2, 4, 1, 5}
	for _, key := range keys {
		tree.Add(key, key*key)
	}
	tree.DisplayInOrder()
	tree.Remove(2)
	tree.Update(5, 6, 6*6)
	tree.DisplayInOrder()
}
