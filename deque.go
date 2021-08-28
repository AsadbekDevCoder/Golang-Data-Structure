package main

import "fmt"

const MAX int = 10

func addfront(arr *[MAX]int, item int, pfront, prear *int) {

	if *pfront == 0 && *prear == MAX-1 {
		fmt.Println("Deque is FULL")
	}

	if *pfront == -1 {
		*pfront = 0
		*prear = 0
		arr[*pfront] = item
		return
	}

	if *prear != MAX-1 {
		c := count(arr)
		k := *prear + 1
		for i := 1; i <= c; i++ {
			arr[k] = arr[k-1]
			k--
		}
		arr[k] = item
		*pfront = k
		*prear++
	}
}

func addrear(arr *[MAX]int, item int, pfront, prear *int) {

	if *pfront == 0 && *prear == MAX-1 {
		fmt.Println("Deque is Full")
	}

	if *pfront == -1 {
		*prear = 0
		*pfront = 0
		arr[*prear] = item
	}

	if *pfront == MAX-1 {
		k := *pfront - 1
		for i := *pfront; i < *prear; i++ {
			k = i
			if k == MAX-1 {
				arr[k] = 0
			} else {
				arr[k] = arr[i+1]
			}
		}
		*prear--
		*pfront--
	}
	*prear++
	arr[*prear] = item
}

func delfront(arr *[MAX]int, pfront, prear *int) int {

	if *pfront == -1 {
		fmt.Println("Deque is Empty !!")
		return 0
	}

	item := arr[*pfront]
	arr[*pfront] = 0

	if *pfront == *prear {
		*pfront, *prear = -1, -1
	} else {
		(*pfront)++
	}

	return item
}

func delrear(arr *[MAX]int, pfront, prear *int) int {

	if *pfront == -1 {
		fmt.Println("Deque is Empty !!")
		return 0
	}

	item := arr[*prear]
	arr[*prear] = 0
	(*prear)--

	if *prear == -1 {
		*pfront = -1
	}

	return item
}

func count(arr *[MAX]int) int {

	c := 0

	for i := 0; i < MAX; i++ {

		if arr[i] != 0 {
			c++
		}
	}

	return c
}

func display(arr *[MAX]int) {

	fmt.Println("Deque element:")
	for _, i := range arr {

		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func main() {

	var arr [MAX]int
	var front, rear = -1, -1

	addfront(&arr, 4, &front, &rear)
	addrear(&arr, 12, &front, &rear)
	addfront(&arr, 5, &front, &rear)
	addrear(&arr, 10, &front, &rear)
	addfront(&arr, 6, &front, &rear)
	addfront(&arr, 7, &front, &rear)
	addrear(&arr, 9, &front, &rear)
	display(&arr)

	delfront(&arr, &front, &rear)
	delfront(&arr, &front, &rear)
	delfront(&arr, &front, &rear)
	addfront(&arr, 7, &front, &rear)
	delrear(&arr, &front, &rear)

	display(&arr)

	fmt.Println("Element count:", count(&arr))
}
