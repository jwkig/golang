package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {
}

type Aircraft struct {
}

type Stream interface {
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}

func closeStream(stream Stream) {
	stream.close()
}

type File struct {
	text string
}

type Folder struct{}

func (f *File) read() string {
	return f.text
}

func (f *File) write(s string) {
	f.text = s
	fmt.Println("Запись в файл строки", s)
}

func (f *File) close() {
	fmt.Println("Файл закрыт")
}

func (f *Folder) close() {
	fmt.Println("Папка закрыта")
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

	file := &File{}
	folder := &Folder{}

	writeToStream(file, "Hello world")
	closeStream(file)

	folder.close()
}
