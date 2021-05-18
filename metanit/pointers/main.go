package main

import "fmt"

func main() {
	var a int = 4
	var p *int = &a
	*p = 100
	fmt.Println("Address: ", p)
	fmt.Println("Value: ", a)
	p = new(int)
	*p = 500
	fmt.Println("Address: ", p)
	fmt.Println("Value: ", *p)

	d := 5
	fmt.Println("d before:", d)
	changeValue(&d)
	fmt.Println("d after:", d)

	p1 := createPointer(7)
	fmt.Println("p1:", *p1)
	p2 := createPointer(10)
	fmt.Println("p2:", *p2)
	p3 := createPointer(28)
	fmt.Println("p3:", *p3)
}

func changeValue(x *int) {
	*x = (*x) * (*x)
}

func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}
