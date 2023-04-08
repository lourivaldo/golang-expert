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

type Person interface {
	Inactive()
}

func (c Client) Inactive() {
	c.Active = false
}

func Inactivation(person Person) {
	person.Inactive()
}

func main() {
	loro := Client{
		Name:   "Loro",
		Age:    28,
		Active: true,
	}
	Inactivation(loro)
	fmt.Println(loro.Active)
}
