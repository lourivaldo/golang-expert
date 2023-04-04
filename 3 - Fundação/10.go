package main

import (
	"fmt"
)

func main() {
	func() {
		// rodar webserver
	}()

	total := func() int {
		return sum(10, 20, 50) * 2
	}()

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
