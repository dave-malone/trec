package trec

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//this implementation is still untested
type amazonSesEmailSender struct {
	endpoint, accessKeyID, secretAccessKey string
}

func newAmazonSesEmailSender(endpoint, accessKeyID, secretAccessKey string) (emailSender amazonSesEmailSender) {
	return amazonSesEmailSender{
		endpoint:        endpoint,
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
	}
}

func (sender *amazonSesEmailSender) send(email emailMessage) error {
	data := make(url.Values)
	data.Add("Action", "SendEmail")
	data.Add("Source", email.from)
	data.Add("Destination.ToAddresses.member.1", email.to)
	data.Add("Message.Subject.Data", email.subject)
	data.Add("Message.Body.Text.Data", email.bodyText)
	data.Add("Message.Body.Html.Data", email.bodyHTML)
	data.Add("AWSAccessKeyId", sender.accessKeyID)

	responseBody, err := sender.sesPost(data)

	fmt.Printf("send email ses response body: %v\n", responseBody)

	return err
}

func (sender *amazonSesEmailSender) setAuthorizationHeader(req *http.Request) {
	now := time.Now().UTC()
	// date format: "Tue, 25 May 2010 21:20:27 +0000"
	date := now.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	req.Header.Set("Date", date)

	h := hmac.New(sha256.New, []uint8(sender.secretAccessKey))
	h.Write([]uint8(date))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	auth := fmt.Sprintf("AWS3-HTTPS AWSAccessKeyId=%s, Algorithm=HmacSHA256, Signature=%s", sender.accessKeyID, signature)
	req.Header.Set("X-Amzn-Authorization", auth)
}

func (sender *amazonSesEmailSender) sesPost(data url.Values) (string, error) {
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", sender.endpoint, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	sender.setAuthorizationHeader(req)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("http error: %s", err)
		return "", err
	}

	resultbody, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if r.StatusCode != 200 {
		log.Printf("error, status = %d", r.StatusCode)

		log.Printf("error response: %s", resultbody)
		return "", fmt.Errorf("error code %d. response: %s", r.StatusCode, resultbody)
	}

	return string(resultbody), nil
}
