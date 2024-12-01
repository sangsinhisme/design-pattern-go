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

type SlackNotifier struct{}

func (s SlackNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender Slack)\n", message)
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

type NotifyDecorator struct {
	core     *NotifyDecorator
	notifier Notifier
}

func (nd NotifyDecorator) Send(message string) {
	nd.notifier.Send(message)
	if nd.core != nil {
		nd.core.Send(message)
	}
}

func (nd NotifyDecorator) Decorator(notifier Notifier) NotifyDecorator {
	return NotifyDecorator{
		core:     &nd,
		notifier: notifier,
	}
}

func CreateNotifyDecorator(notifier Notifier) NotifyDecorator {
	return NotifyDecorator{notifier: notifier}
}

func main() {
	notifier := CreateNotifyDecorator(EmailNotifier{}).
		Decorator(SMSNotifier{}).
		Decorator(SlackNotifier{})

	service := NotificationService{
		notifier: notifier,
	}

	service.SendNotification("Hello world!")
}
