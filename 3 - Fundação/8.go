package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(10, 20))

	value, err := sum(30, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

// func sum(a int, b int) int {
// func sum(a, b int) (int, bool) {
// 	if a+b >= 50 {
// 		return a + b, true
// 	}
// 	return a + b, false
// }

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
