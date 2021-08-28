package main

import "fmt"

const size int = 5

var items [size]int
var front, rear = -1, -1

func isfull() bool {

	if (front == rear+1) || (front == 0 && rear == size-1) {

		return true
	}

	return false
}

func enqueue(data int) {

	if isfull() {
		fmt.Println("Circular Queue is Full!!")
	} else {

		if front == -1 {
			front = 0
		}
		rear = (rear + 1) % size
		items[rear] = data
		fmt.Println("Inserted element: ", data)
	}
}

func isempty() bool {

	if front == -1 {

		return true
	}

	return false
}

func dequeue() {

	if isempty() {
		fmt.Println("Circular Queue is Empty!!")
	} else {

		element := items[front]
		if front == rear {
			front, rear = -1, -1
		} else {
			front = (front + 1) % size
		}

		fmt.Println("Deleted element:", element)
	}
}

func display() {

	var i int
	if isempty() {
		fmt.Println("Circular Queue is Empty !!")
	} else {
		fmt.Println("Circular Queue:")
		for i = front; i != rear; i = (i + 1) % size {
			fmt.Printf("%d ", items[i])
		}
		fmt.Printf("%d \n", items[i])
	}
}

func main() {

	dequeue()

	enqueue(1)
	enqueue(2)
	enqueue(3)
	enqueue(4)
	enqueue(5)

	display()

	dequeue()

	display()
}
