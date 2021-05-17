package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, "Иван", "Иванов"))
}

func add(x, y int, firstName, lastName string) (sum int, fullName string) {
	sum = x + y
	fullName = firstName + " " + lastName
	return
}
