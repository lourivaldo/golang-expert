package main

import "fmt"

type Client struct {
	name string
}

//	func (c Client) walk() {
//		c.name = "Lourivaldo"
//		fmt.Printf("Client %v walk\n", c.name)
//	}
func (c *Client) walk() {
	c.name = "Lourivaldo"
	fmt.Printf("Client %v walk\n", c.name)
}

func main() {
	loro := Client{
		name: "Jhon",
	}
	loro.walk()
	fmt.Printf("Value of struct is %v\n", loro.name)
}
