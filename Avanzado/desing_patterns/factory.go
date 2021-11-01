package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}
type Computer struct {
	stock int
	name  string
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}
func (c *Computer) setName(name string) {
	c.name = name
}
func (c *Computer) getStock() int {
	return c.stock
}
func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{Computer: Computer{stock: 25, name: "Laptop Computer"}}
}

type Desktop struct {
	Computer
}

func newDesktop() IProduct {
	return &Desktop{Computer: Computer{stock: 50, name: "Desktop Computer"}}
}
func GetComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "Laptop":
		return newLaptop(), nil
	case "Desktop":
		return newDesktop(), nil
	default:
		return nil, fmt.Errorf("invalid computer type")
	}
}

func printNameAndStock(p IProduct) {
	fmt.Printf("Name: %s, Stock: %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("Laptop")
	desktop, _ := GetComputerFactory("Desktop")
	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
