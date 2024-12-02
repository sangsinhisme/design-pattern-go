package main

import "fmt"

type Doctor struct {
	next Department
}

func (r *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Doctor) setNext(next Department) {
	r.next = next
}
