package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	var numbers = [...]int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		fmt.Println(num)
	}

	var users = [3]string{"Tom", "Alice", "Kate"}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
