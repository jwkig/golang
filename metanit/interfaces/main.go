package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {
}

type Aircraft struct {
}

func (car Car) move() {
	fmt.Println("Car moving")
}

func (plain Aircraft) move() {
	fmt.Println("Aircraft flying")
}

func main() {
	var car Car = Car{}
	var plane Aircraft = Aircraft{}

	car.move()
	plane.move()
}
