package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Email struct {
	Sender, Subject, Body string
	Recipients            []string
	Attachments           []os.File
}

type EmailBuilder struct {
	email *Email
}

func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{email: &Email{}}
}

func (eb *EmailBuilder) Sender(sender string) *EmailBuilder {
	if !strings.Contains(sender, "@") {
		fmt.Println("Warning: Invalid email format")
		return eb
	}
	eb.email.Sender = sender
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.Subject = subject
	return eb
}

func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.Body = body
	return eb
}

func (eb *EmailBuilder) AddRecipient(recipient string) *EmailBuilder {
	if !strings.Contains(recipient, "@") {
		fmt.Println("Warning: Invalid recipient email")
		return eb
	}
	eb.email.Recipients = append(eb.email.Recipients, recipient)
	return eb
}

func (eb *EmailBuilder) Attach(attachment *os.File) *EmailBuilder {
	if attachment == nil {
		panic("Attachment cannot be nil")
	}
	eb.email.Attachments = append(eb.email.Attachments, *attachment)
	return eb
}

// Validation before sending
func (eb *EmailBuilder) Validate() error {
	if eb.email.Sender == "" {
		return errors.New("Sender email is required")
	}
	if len(eb.email.Recipients) == 0 {
		return errors.New("At least one recipient is required")
	}
	return nil
}

func sendEmailImpl(email *Email) {
	fmt.Println("Sending email to", email.Recipients)
}

type builder func(*EmailBuilder)

func SendEmail(action builder) error {
	builder := EmailBuilder{email: &Email{}}
	action(&builder)

	// Validate email before sending
	if err := builder.Validate(); err != nil {
		return err
	}

	sendEmailImpl(builder.email)
	return nil
}

// **Main Function**
func main() {
	err := SendEmail(func(b *EmailBuilder) {
		b.Sender("sinhns2@fpt.com").
			AddRecipient("alice@gmail.com").
			Subject("Hello").
			Body("This is a test email.")
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
