package trec

import "fmt"

type (
	emailSender interface {
		send(email emailMessage) error
	}

	emailSenderFactory func() emailSender
)

var (
	newEmailSender emailSenderFactory
)

type emailMessage struct {
	from, to, subject, bodyText, bodyHTML string
}

func newEmailMessage(from, to, subject, bodyText, bodyHTML string) *emailMessage {
	return &emailMessage{
		from:     from,
		to:       to,
		subject:  subject,
		bodyText: bodyText,
		bodyHTML: bodyHTML,
	}
}

func (e *emailMessage) String() string {
	return fmt.Sprintf("emailMessage{from:%v, to:%v, subject:%v}", e.from, e.to, e.subject)
}
