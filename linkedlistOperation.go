package main

import (
	"errors"
	"fmt"
)

// Node represents a node of linked list
type Node struct {
	data int
	next *Node
}

// Linked represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(data int) {

	newNode := Node{}
	newNode.data = data

	if l.len == 0 {

		l.head = &newNode
		l.len++
		return
	}

	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &newNode
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) InsertAt(pos int, data int) {

	// Create new node
	newNode := Node{}
	newNode.data = data

	// Validate the position
	if pos < 0 {
		return
	}

	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}

	if pos > l.len {
		return
	}

	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

func (l *LinkedList) GetAt(pos int) *Node {

	ptr := l.head
	if pos < 0 {
		return ptr
	}

	if pos > (l.len - 1) {
		return nil
	}

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr
}

func (l *LinkedList) Search(data int) bool {

	ptr := l.head
	for i := 0; i < l.len; i++ {

		if ptr.data == data {
			return true
		}
		ptr = ptr.next
	}

	return false
}

func (l *LinkedList) DeleteAt(pos int) error {

	// validate the position
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in List")
	}

	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}

	prevNode.next = l.GetAt(pos).next
	l.len--
	return nil
}

func (l *LinkedList) Deletedata(data int) error {

	ptr := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}

	for i := 0; i < l.len; i++ {

		if ptr.data == data {

			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}

			l.len--
			return nil
		}

		ptr = ptr.next
	}

	fmt.Println("\nNode not found")
	return errors.New("Node not found")
}

func (l *LinkedList) Sorting() {

	current := l.head

	if l.len == 0 {
		fmt.Println("Not found Node")
		return
	} else {
		for current != nil {

			index := current.next

			for index != nil {

				if current.data > index.data {

					temp := current.data
					current.data = index.data
					index.data = temp
				}

				index = index.next
			}

			current = current.next
		}
	}
}

func (l *LinkedList) Print() {

	if l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}

	ptr := l.head

	fmt.Println("\nLinked List")
	for i := 0; i < l.len; i++ {
		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}

	fmt.Println()
}

func main() {

	llist := LinkedList{}

	llist.Print()

	llist.Insert(3)
	llist.Insert(5)
	llist.Insert(7)
	llist.Insert(8)

	llist.InsertAt(2, 9)
	llist.InsertAt(1, 11)

	llist.Print()

	search := 9

	if llist.Search(search) {
		fmt.Println("\nData", search, "is found")
	} else {
		fmt.Println("\nData", search, "is not found")
	}

	llist.DeleteAt(3)
	llist.Print()

	llist.Deletedata(33)
	llist.Print()

	llist.Sorting()
	llist.Print()
}
