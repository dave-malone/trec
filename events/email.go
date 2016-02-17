package events

import (
	"fmt"

	"github.com/dave-malone/email"
)

//EmailEvent an Event that should result in the sending of an email.Message
type EmailEvent struct {
	*appEvent
	Email *email.Message
}

type EmailEventObserver struct {
}

func (o *EmailEventObserver) Receive(e Event) {
	if emailEvent, ok := e.(EmailEvent); ok {
		in := pub(emailEvent.Email)
		sub(in)
	}
}

func pub(m *email.Message) <-chan *email.Message {
	emailChan := make(chan *email.Message)

	go func() {
		emailChan <- m
		close(emailChan)
	}()

	return emailChan
}

func sub(in <-chan *email.Message) {
	for message := range in {
		fmt.Printf("Received email message: %v", *message)
		sender := email.NewSender()
		sender.Send(message)
	}
}
