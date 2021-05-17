package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, "Иван", "Иванов"))

	var funcPtr func(int, int) int = simpleAdd

	fmt.Println("Additions")
	printFunc(funcPtr)
	fmt.Println("Multiplications")
	printFunc(multiply)

}

func add(x, y int, firstName, lastName string) (sum int, fullName string) {
	sum = x + y
	fullName = firstName + " " + lastName
	return
}

func printFunc(funcPtr func(int, int) int) {
	var numbers = [...]int{1, 2, 3}
	for _, number := range numbers {
		fmt.Println(funcPtr(number, number))
	}
}
func simpleAdd(x, y int) int { return x + y }
func multiply(x, y int) int  { return x * y }
