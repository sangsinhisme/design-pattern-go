package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"}}

	//jane := john

	// shallow copy-method
	//jane.Name = "Jane" // ok

	//jane.Address.StreetAddress = "321 Baker St"

	//fmt.Println(john.Name, john.Address)
	//fmt.Println(jane.Name, jane. Address)

	// what you really want
	// deep copy-method

	// However, the problem with this approach is that it really doesn't scale.
	// If you have a person which is composed of address and address is itself composed of something else,
	// then you end up with this very complicated recursive structure and you end up having to write lots of
	// code just to be able to copy-method objects.
	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country}

	jane.Name = "Jane" // ok
	jane.Address.StreetAddress = "321 Baker St"

	fmt.Println(john.Name, john.Address)
	fmt.Println(jane.Name, jane.Address)
}
