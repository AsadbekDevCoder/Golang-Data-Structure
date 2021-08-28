package main

import "fmt"

func heapify(slice []int, n, i int) []int {

	if n == 1 {

		fmt.Println("Single element")
	} else {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && slice[left] > slice[largest] {

			largest = left
		}

		if right < n && slice[right] > slice[largest] {

			largest = right
		}

		if largest != i {
			slice[i], slice[largest] = slice[largest], slice[i]
			heapify(slice, n, largest)
		}
	}

	return slice
}

func insert(slice []int, newItem int) []int {

	size := len(slice)
	if size == 0 {
		slice = append(slice, newItem)
	} else {
		slice = append(slice, newItem)
		for i := size/2 - 1; i > -1; i-- {

			slice = heapify(slice, size, i)
		}
	}

	return slice
}

func main() {

	var slice = []int{}

	slice = insert(slice, 3)
	slice = insert(slice, 5)
	slice = insert(slice, 1)
	slice = insert(slice, 10)
	slice = insert(slice, 13)

	fmt.Println(slice)

}
