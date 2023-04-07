package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	loro := Client{
		Name:   "Loro",
		Age:    28,
		Active: true,
	}
	loro.Active = false

	fmt.Printf("Name: %s, Age: %s, Active: %s", loro.Name, loro.Age, loro.Active)
}
