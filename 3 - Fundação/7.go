package main

import "fmt"

func main() {
	salarios := map[string]int{"Loro": 1000, "Isla": 500}
	fmt.Println(salarios["Isla"])
	fmt.Println(salarios)

	delete(salarios, "Isla")
	fmt.Println(salarios)

	salarios["Loro"] = 5000
	delete(salarios, "Isla")
	fmt.Println(salarios)

	sal :=
		make(map[string]int)
	fmt.Println(sal)
	sal["Pedro"] = 9000
	fmt.Println(sal)

	for name, salario := range salarios {
		fmt.Printf("O salario de %s e %d\n", name, salario)
	}
	for _, salario := range salarios {
		fmt.Printf("O salario e %d\n", salario)
	}
}
