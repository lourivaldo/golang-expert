package main

func main() {
	// Memoria -> Endereço -> Valor

	// variavel -> ponteiro que tem um endereço na memoria -> ten valor
	a := 10
	var pointer *int = &a
	*pointer = 20
	b := &a
	println(pointer)
	println(a)
	println(b)
	println(*b)
	*b = 30
	println(*b)
	println(a)
}
