package trec

import "github.com/xchapter7x/lo"

type noopEmailSender struct {
	messages []emailMessage
}

func newNoopEmailSender() noopEmailSender {
	return noopEmailSender{
		messages: []emailMessage{},
	}
}

func (sender noopEmailSender) send(email emailMessage) error {
	lo.G.Debug("noopEmailSender.send(%v)\n", email)
	sender.messages = append(sender.messages, email)

	return nil
}
