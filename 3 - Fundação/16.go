package main

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	minVar1 := 10
	minVar2 := 20
	println(sum(&minVar1, &minVar2))
	println(minVar1)
}
