package main

import "fmt"

type library []string

func (l library) print() {

	for _, val := range l {
		fmt.Println(val)
	}
}

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}

func main() {
	var lib library = library{"Book1", "book2", "book3"}
	lib.print()

	var tom = person{name: "Tom", age: 24}
	tom.print()
	tom.eat("борщ с капустой, но не красный")
}
