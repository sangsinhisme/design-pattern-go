package main

import "fmt"

type NotificationService struct {
	notifierType string
}

/*
switch case
if else same function
*/

func (n NotificationService) SendNotification(message string) {
	if n.notifierType == "email" {
		fmt.Printf("Sending message: %s (Sender Email)\n", message)
	} else if n.notifierType == "sms" {
		fmt.Printf("Sending message: %s (Sender SMS)\n", message)
	}
}

func main() {
	s := NotificationService{notifierType: "email"}
	s.SendNotification("Hello world!")
}
