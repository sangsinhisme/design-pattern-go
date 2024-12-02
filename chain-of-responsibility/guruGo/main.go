package main

func main() {
	cashier := &Cashier{}
	medical := &Medical{next: cashier}
	doctor := &Doctor{next: medical}
	reception := &Reception{next: doctor}
	patient := &Patient{name: "em TungPVP"}
	reception.execute(patient)
}
