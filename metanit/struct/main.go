package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name        string
	age         int
	contactInfo contact
}

func main() {
	var tom = person{
		name: "Tom",
		age:  24,
		contactInfo: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	fmt.Println(tom.name)
	fmt.Println(tom.age)

	tom.age = 38
	fmt.Println(tom.name, tom.age)

	var tomPointer *person = &tom
	tomPointer.age = 29
	fmt.Println(tom.age)
	(*tomPointer).age = 32
	fmt.Println(tom.age)
}
