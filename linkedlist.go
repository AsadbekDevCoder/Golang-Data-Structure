package main

import "fmt"

type node struct {
	item int
	next *node
}

func display(linked *node) {

	for linked != nil {
		fmt.Printf("%d ", linked.item)
		linked = linked.next
	}
}

func main() {

	var first node = node{item: 3}
	var second node = node{item: 4}
	var third node = node{item: 6}

	first.next = &second
	second.next = &third

	display(&first)
}
