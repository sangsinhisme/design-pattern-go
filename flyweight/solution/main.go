package main

import "fmt"

type ChatMessage struct {
	Content string
	Sender  *Sender
}

type Sender struct {
	Name   string
	Avatar []byte
}

type SenderFactory struct {
	cacheSender map[string]*Sender
}

func (sf SenderFactory) GetSender(name string) *Sender {
	return sf.cacheSender[name]
}

func main() {
	senderFactory := SenderFactory{
		cacheSender: map[string]*Sender{
			"Peter": {
				Name:   "Peter",
				Avatar: make([]byte, 1024*300),
			},
			"Mary": {
				Name:   "Mary",
				Avatar: make([]byte, 1024*400),
			},
		},
	}
	fmt.Println([]ChatMessage{
		{
			Content: "hi",
			Sender:  senderFactory.GetSender("Peter"),
		},
		{
			Content: "hello Peter",
			Sender:  senderFactory.GetSender("Mary"),
		},
		{
			Content: "how are you?",
			Sender:  senderFactory.GetSender("Peter"),
		},
		{
			Content: "bien!",
			Sender:  senderFactory.GetSender("Mary"),
		},
	})
}
