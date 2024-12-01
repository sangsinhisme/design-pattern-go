package main

import "fmt"

type ChatMessage struct {
	Content      string
	SenderName   string
	SenderAvatar []byte
}

func main() {
	fmt.Println([]ChatMessage{
		{
			Content:      "hi",
			SenderName:   "Peter",
			SenderAvatar: make([]byte, 1024*300),
		},
		{
			Content:      "hi Peter",
			SenderName:   "Mary",
			SenderAvatar: make([]byte, 1024*400),
		},
		{
			Content:      "how are you?",
			SenderName:   "Peter",
			SenderAvatar: make([]byte, 1024*300),
		},
		{
			Content:      "bien!",
			SenderName:   "Mary",
			SenderAvatar: make([]byte, 1024*400),
		},
	})
}
