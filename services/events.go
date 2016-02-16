package trec

import (
	"fmt"

	"github.com/dave-malone/email"
	"github.com/xchapter7x/lo"
)

func newUserRegistrationEvent(user User) {
	lo.G.Debugf("user registration event for user %v", user)

	emailMessage := email.NewMessage(
		"no-reply@therealestatecrm.com",
		user.Email,
		"TheRealEstateCRM.com Verification",
		"{bodyText}",
		"{bodyHtml}",
	)

	in := sendEmailEvent(emailMessage)
	emailEventObserver(in)
}

func sendEmailEvent(message *email.Message) <-chan *email.Message {
	emailChan := make(chan *email.Message)

	go func() {
		emailChan <- message
		close(emailChan)
	}()

	return emailChan
}

func emailEventObserver(in <-chan *email.Message) {
	for message := range in {
		fmt.Printf("Received email message: %v", *message)
		sender := email.NewSenderFactory()
		sender.Send(message)
	}

}
