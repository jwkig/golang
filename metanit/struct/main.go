package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name string
	age  int
	contact
}

type node struct {
	value int
	next  *node
}

func printNodeValue(n *node) {

	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

func main() {
	var tom = person{
		name: "Tom",
		age:  24,
		contact: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	fmt.Println(tom.name)
	fmt.Println(tom.age)
	fmt.Println(tom.email)

	tom.age = 38
	fmt.Println(tom.name, tom.age)

	var tomPointer *person = &tom
	tomPointer.age = 29
	fmt.Println(tom.age)
	(*tomPointer).age = 32
	fmt.Println(tom.age)

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}

}
