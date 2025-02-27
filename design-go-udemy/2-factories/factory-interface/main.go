package main

import "fmt"

type Person interface {
	SayHello()
}
type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hello, my name is", p.name)
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Hello, my name is", p.name, "and I am tired")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("John", 30)
	//p.age++ -> can not access the age field
	p.SayHello()

	// -> Different time of object and return only interface Person
	p2 := NewPerson("Jane", 101)
	p2.SayHello()

}
