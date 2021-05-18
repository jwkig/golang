package main

import "fmt"

func main() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	}
	delete(people, "Bob")
	for key, value := range people {
		fmt.Println(key, value)
	}
}
