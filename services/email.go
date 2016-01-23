package trec

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

type emailSender interface {
	send(email emailMessage) (err error)
}
