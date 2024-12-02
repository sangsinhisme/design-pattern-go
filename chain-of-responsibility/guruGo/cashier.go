package main

import "fmt"

type Cashier struct {
	next Department
}

func (r *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (r *Cashier) setNext(next Department) {
	r.next = next
}
