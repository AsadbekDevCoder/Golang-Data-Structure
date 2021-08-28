package main

import "fmt"

const SIZE uint8 = 5

var items [SIZE]int
var front, rear = -1, -1

func enqueue(item int) {

	if uint8(rear) == SIZE-1 {

		fmt.Println("Queue is Full!!!")
	} else {

		if front == -1 {

			front = 0
		}
		rear++
		items[rear] = item
	}
}

func dequeue() {

	if rear == -1 {

		fmt.Println("Queue is Empty!!!")
	} else {

		fmt.Println("Deleted item:", items[front])
		front++
		if front > rear {

			front, rear = -1, -1
		}
	}
}

func display() {

	if rear == -1 {

		fmt.Println("Queue is Empty!!!")
	} else {
		fmt.Println("Queue Items:")
		for i := rear; i >= front; i-- {

			fmt.Println(items[i])
		}
	}
}

func main() {

	enqueue(2)
	enqueue(3)
	enqueue(4)
	enqueue(5)
	enqueue(6)

	display()

	dequeue()
	dequeue()
	dequeue()

	display()

}
