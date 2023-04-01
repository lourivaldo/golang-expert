package main

import "fmt"

var myArray [3]int

func main() {
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	// myArray[3] = 40

	fmt.Printf("O tipo de myArray é %T\n", myArray)
	fmt.Println(myArray[len(myArray)-1])

	for i, v := range myArray {
		fmt.Printf("O valor é %d e o index é %d\n", v, i)
	}
}
