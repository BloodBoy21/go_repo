package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}
type FullTimeEmployee struct {
	Person
	Employee
}

type TemporatyEmployee struct {
	Person
	Employee
	taxRate int
}

func getMessage(p PrintInfo){
	fmt.Println(p.getMessage())
}

func (ftEmployee FullTimeEmployee)getMessage()string{
	return "FullTimeEmployee"
}
func (tEmployee TemporatyEmployee)getMessage()string{
	return "TemporatyEmployee"
}
type PrintInfo interface  {
	getMessage() string
}


func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}


func GetFullTimeEmployee(id int,dni string) (FullTimeEmployee, error){
	var ftEmployee FullTimeEmployee
	e,err:=GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee =e
	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p
	return ftEmployee, nil



}

func GetPersonByDNI(dni string) (Person, error){
	time.Sleep(time.Second*5)
	return Person{},nil
}
func GetEmployeeById(id int)(Employee,error){
	time.Sleep(time.Second*5)
	return Employee{},nil
}


func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Alan"
	ftEmployee.age = 19
	ftEmployee.id = 321
	fmt.Printf("%v",ftEmployee)
	tEmployee := TemporatyEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee) 
}