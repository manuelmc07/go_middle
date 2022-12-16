package main

import "fmt"

// Person public struct
type Person struct {
	name string
	age  int
}

// Employee public struct
type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full-time employee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

type PrintInfo interface {
	getMessage() string
}

func GetMessage(p Person) {
	fmt.Printf("%s have %d years old", p.name, p.age)
}

func main() {
	// inheritance  composition
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Ramiro"
	ftEmployee.age = 45
	ftEmployee.id = 5
	fmt.Println(ftEmployee)
	tEmployee := TemporaryEmployee{}

	// polymorphism
	GetMessage(ftEmployee.Person)

	getMessage(tEmployee)
	getMessage(ftEmployee)
}
