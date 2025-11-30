package mailer

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/onosejoor/email-api/go-server-without-creds/types"
	gomail "gopkg.in/mail.v2"
)

var (
	dialer *gomail.Dialer
	sender gomail.SendCloser
	once   sync.Once
)

func initDialer(gmailUser, gmailAppPassword string) {
	dialer = gomail.NewDialer("smtp.gmail.com", 465, gmailUser, gmailAppPassword)
	var err error
	sender, err = dialer.Dial()
	if err != nil {
		log.Fatalf("Failed to dial SMTP: %v", err)
	}
}

// SendEmail sends using a persistent SMTP connection
func SendEmail(req types.SendEmailRequest) error {
	once.Do(func() { initDialer(req.GMAIL_USER, req.GMAIL_APP_PASSWORD) })

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", req.From, req.GMAIL_USER))
	m.SetHeader("To", strings.Join(req.To, ","))
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", req.HTML)

	err := gomail.Send(sender, m)
	if err != nil {
		sender, _ = dialer.Dial()
		return gomail.Send(sender, m)
	}
	return nil

}
