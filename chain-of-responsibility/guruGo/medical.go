package main

import "fmt"

type Medical struct {
	next Department
}

func (r *Medical) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Medicine already given to patient")
		r.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Medical) setNext(next Department) {
	r.next = next
}
