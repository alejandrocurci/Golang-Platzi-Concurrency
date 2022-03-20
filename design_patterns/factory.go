package main

import "fmt"

type Product interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) getStock() int {
	return c.stock
}

type Laptop struct {
	Computer
}

func newLaptop() Product {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

func newDesktop() Product {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

func NewComputer(computerType string) (Product, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	}

	if computerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("invalid computer type")
}

func printNameAndStock(p Product) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := NewComputer("laptop")
	desktop, _ := NewComputer("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
