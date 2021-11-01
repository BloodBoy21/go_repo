package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int){
	e.id = id 
}
func (e*Employee) SetName(name string){
	e.name = name
}

func newEmployee(name string,id int)*Employee{
	return &Employee{
		id: id,
		name: name,
	}
}

func main() {
	e := Employee{id: 1, name: "Alan"}
	fmt.Printf("e: %v\n", e)
	e.SetId(10)
	e.SetName("Bob")
	fmt.Printf("e: %v\n",e)
	alan := newEmployee("Alan",122)
	fmt.Printf("Employee: %v\n", *alan)
}