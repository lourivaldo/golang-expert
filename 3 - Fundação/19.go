package main

import "fmt"

func main() {
	var myVar interface{} = "Lourivaldo"

	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
	res2 := myVar.(int)
	fmt.Printf("O valor de res é %v", res2)
}
