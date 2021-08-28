package main

import (
	"fmt"
)

type text string

type Customer struct {
	firstname, lastname text
	age                 uint8
	birth               uint16
	workplace           text
	next                *Customer
}

func printList(user *Customer) {

	for user != nil {

		fmt.Printf(
			"%-10s| %-10s| %-4d| %-5d| %-10s\n",
			user.firstname,
			user.lastname,
			user.age,
			user.birth,
			user.workplace,
		)
		user = user.next
	}
}

func main() {

	var first Customer = Customer{
		firstname: "Asadbek",
		lastname:  "Ergashev",
		age:       21,
		birth:     2000,
		workplace: "Najot Talim",
	}

	var second Customer = Customer{
		firstname: "Najim",
		lastname:  "Sobirov",
		age:       23,
		birth:     1998,
		workplace: "PDP",
	}

	var third Customer = Customer{
		firstname: "Oybek",
		lastname:  "Akmalov",
		age:       18,
		birth:     2003,
		workplace: "School",
	}

	first.next = &second
	second.next = &third

	printList(&first)

}
