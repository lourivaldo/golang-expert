package main

//func Soma(m map[string]int) int {
//	var total int
//	for _, v := range m {
//		total += v
//	}
//	return total
//}

//func Soma[T int | float64](m map[string]T) T {
//	var total T
//	for _, v := range m {
//		total += v
//	}
//	return total
//}

//type Number interface {
//	int | float64
//}

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 1000.20, "João": 2000.10, "Maria": 3000.40}
	m3 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	//println(Compara(10, 10.0))
	println(Compara(10, 10))
}
