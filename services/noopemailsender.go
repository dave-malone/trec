package trec

import "github.com/xchapter7x/lo"

type noopEmailSender struct {
	messages []emailMessage
}

func newNoopEmailSender() emailSender {
	return noopEmailSender{
		messages: []emailMessage{},
	}
}

func (sender noopEmailSender) send(email emailMessage) error {
	lo.G.Debugf("email received: %v\n", email)
	sender.messages = append(sender.messages, email)

	return nil
}
