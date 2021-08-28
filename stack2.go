package main

import "fmt"

func createStack() []int {

	var stack []int
	return stack
}

func check_empty(stack []int) bool {

	return len(stack) == 0
}

func push(stack *[]int, item int) {

	*stack = append(*stack, item)
	fmt.Println("Push item:", item)
}

func pop(stack []int) []int {

	if check_empty(stack) {

		fmt.Println("Stack is empty")
		return nil
	}

	fmt.Println("Stack popped item: ", stack[len(stack)-1])
	stack = append(stack[:len(stack)-1], stack[len(stack):]...)
	return stack
}

func main() {

	stack := createStack()

	push(&stack, 3)
	push(&stack, 4)
	push(&stack, 6)

	fmt.Println(stack)

	stack = pop(stack)

	fmt.Println(stack)
}
