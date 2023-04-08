package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

func main() {
	loro := Client{
		Name:   "Loro",
		Age:    28,
		Active: true,
	}
	loro.Active = false

	//loro.City = "Recife"
	loro.Address.City = "Recife"

	fmt.Printf("Name: %s, Age: %s, Active: %s", loro.Name, loro.Age, loro.Active)
}
