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

func (c Client) Inactive() {
	c.Active = false
}

func main() {
	loro := Client{
		Name:   "Loro",
		Age:    28,
		Active: true,
	}
	loro.Inactive()
	fmt.Println(loro.Active)
}
