package main

import "fmt"

// Employee public struct
type Employee struct {
	id   int
	name string
}

// SetId public function
func (e *Employee) SetId(id int) {
	e.id = id
}

// SetName public function
func (e *Employee) SetName(name string) {
	e.name = name
}

// GetId public function
func (e *Employee) GetId() int {
	return e.id
}

// GetName public function
func (e *Employee) GetName() string {
	return e.name
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func main() {
	// declare new instance of employee struct
	e := Employee{}
	// add setters
	e.SetId(5)
	e.SetName("Manuel")
	// retrieve getters
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

	// 2
	e2 := Employee{
		id:   1,
		name: "Luis",
	}
	fmt.Println(e2)

	// 3
	// whe used new keyword to instance will be work with pointer
	e3 := new(Employee)
	e3.id = 4
	e3.name = "manuel"
	fmt.Println(*e3)

	// 4
	e4 := Employee{}
	e4.id = 2
	e4.name = "Joss"
	fmt.Println(e4)

	// 5
	e5 := NewEmployee(8, "Vivika")
	fmt.Println(*e5)

}
