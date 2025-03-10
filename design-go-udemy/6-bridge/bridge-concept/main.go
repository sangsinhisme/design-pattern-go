package main

import "fmt"

// MessageSender Implementation Interface: Defines contract for different senders
type MessageSender interface {
	SendMessage(text string)
}

type EmailSender struct{}

func (e *EmailSender) SendMessage(text string) {
	fmt.Println("Sending email: ", text)
}

type SMSSender struct{}

func (s *SMSSender) SendMessage(text string) {
	fmt.Println("Sending sms: ", text)
}

type Notification struct {
	sender MessageSender
}

func (n *Notification) Notify(text string) {
	n.sender.SendMessage(text)
}

type UrgentNotification struct {
	Notification
}

func (u *UrgentNotification) Notify(text string) {
	fmt.Println("[Urgent] Notification Sent:")
	u.sender.SendMessage(text)
}

func main() {
	// Using abstraction with different implementations
	emailNotif := &Notification{sender: &EmailSender{}}
	emailNotif.Notify("Hello via Email")

	smsNotif := &Notification{sender: &SMSSender{}}
	smsNotif.Notify("Hello via SMS")

	urgentSmsNotif := &UrgentNotification{Notification{sender: &SMSSender{}}}
	urgentSmsNotif.Notify("Critical Alert!")
}
