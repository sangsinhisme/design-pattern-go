package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender Email)\n", message)
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender SMS)\n", message)
}

type NotificationService struct {
	notifier Notifier
}

func (s *NotificationService) SetNotifier(notifier Notifier) {
	s.notifier = notifier
}

func (s *NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func CreateNotifier(t string) Notifier {
	if t == "sms" {
		return SMSNotifier{}
	}
	return EmailNotifier{}
}

func main() {
	s := &NotificationService{
		notifier: CreateNotifier("email"),
	}
	s.SendNotification("Hello World!")
	s.SetNotifier(CreateNotifier("sms"))

	s.SendNotification("Hello World!")
}