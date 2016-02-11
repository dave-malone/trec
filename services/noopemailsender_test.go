package trec

import "testing"

func TestNewNoopEmailSender(t *testing.T) {
	emailSender := newNoopEmailSender()

	if emailSender == nil {
		t.Fatal("newNoopEmailSender returned a nil value")
	}
}
