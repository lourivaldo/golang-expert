package main

import "fmt"

const a = "Hello"

type ID int

var id ID = 1

func main() {
	fmt.Printf("O tipo de a é %T\n", a)
	fmt.Printf("O valor de a é %v\n", a)

	fmt.Printf("O tipo de id é %T\n", id)
}
