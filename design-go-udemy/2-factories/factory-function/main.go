package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

/*
Factory function is a function that returns an instance of a struct.
Factory functions are used to create instances of structs that are not directly accessible.
This is the gist of factory functions -> factory functions all over the place when there is some logic that needs to be applied to create an instance of a struct.
*/
func NewPerson(name string, age int) *Person {
	if age < 16 {
		// ...
	}
	return &Person{name, age, 2}
}
func main() {
	p := NewPerson("John", 30)
	fmt.Println(p)
}
