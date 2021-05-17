package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}

	for _, num := range numbers {
		fmt.Println(num)
	}

	var users = [3]string{"Tom", "Alice", "Kate"}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	var sum = 0

	for _, value := range numbers {
		if value < 0 {
			continue // переходим к следующей итерации
		}
		sum += value
	}
	fmt.Println("Sum:", sum)

	numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum = 0

	for _, value := range numbers {
		if value > 4 {
			break // если число больше 4 выходим из цикла
		}
		sum += value
	}
	fmt.Println("Sum:", sum) // Sum: 10
}
