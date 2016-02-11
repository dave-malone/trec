package trec

import (
	"fmt"

	"github.com/xchapter7x/lo"
)

func newUserRegistrationEvent(user User) {
	lo.G.Debugf("user registration event for user %v", user)

	emailMessage := newEmailMessage(
		"no-reply@therealestatecrm.com",
		user.Email,
		"TheRealEstateCRM.com Verification",
		"{bodyText}",
		"{bodyHtml}",
	)

	in := sendEmailEvent(emailMessage)
	emailEventObserver(in)
}

func sendEmailEvent(message *emailMessage) <-chan *emailMessage {
	emailChan := make(chan *emailMessage)

	go func() {
		emailChan <- message
		close(emailChan)
	}()

	return emailChan
}

func emailEventObserver(in <-chan *emailMessage) {
	for message := range in {
		fmt.Printf("Received email message: %v", *message)
		sender := newEmailSender()
		sender.send(*message)
	}

}
