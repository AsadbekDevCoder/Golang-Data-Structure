package main

import "fmt"

const MAX int = 10

var count int = 0

type Stack struct {
	item [MAX]int
	top  int
}

type st *Stack

func createstack(stack st) {

	stack.top = -1
}

func isfull(stack st) bool {

	if stack.top == MAX-1 {

		return true
	}

	return false
}

func push(stack st, newitem int) {

	if isfull(stack) {

		fmt.Println("Stack FULL")
	} else {

		stack.top++
		stack.item[stack.top] = newitem
	}
	count++
}

func isempty(stack st) bool {

	if stack.top == -1 {

		return true
	}

	return false
}

func pop(stack st) {

	if isempty(stack) {

		fmt.Println("Stack is Empty")
	} else {

		fmt.Println("Item popped =", stack.item[stack.top])
		stack.top--
	}
	count--
}

func printList(stack st) {

	for i := 0; i < count; i++ {

		fmt.Printf("%d ", stack.item[i])
	}

	fmt.Println()
}

func main() {

	var stack Stack

	createstack(&stack)

	push(&stack, 3)
	push(&stack, 5)
	push(&stack, 7)
	push(&stack, 9)

	fmt.Println("Stack items:")
	printList(&stack)

	pop(&stack)

	printList(&stack)
}
