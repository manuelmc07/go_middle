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

func main() {
	// declare new instance of employee struct
	e := Employee{}
	// add setters
	e.SetId(5)
	e.SetName("Manuel")
	// retrieve getters
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
