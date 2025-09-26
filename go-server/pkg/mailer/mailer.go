package mailer

import (
	"fmt"
	"log"
	"os"
	"sync"

	gomail "gopkg.in/mail.v2"
)

var (
	dialer *gomail.Dialer
	sender gomail.SendCloser
	once   sync.Once
)

func initDialer() {
	dialer = gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("GMAIL_USER"), os.Getenv("GMAIL_APP_PASSWORD"))
	var err error
	sender, err = dialer.Dial()
	if err != nil {
		log.Fatalf("Failed to dial SMTP: %v", err)
	}
}

// SendEmail sends using a persistent SMTP connection
func SendEmail(from, to, subject, body string) error {
	once.Do(initDialer)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", from, os.Getenv("GMAIL_USER")))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	err := gomail.Send(sender, m)
	if err != nil {
		sender, _ = dialer.Dial()
		return gomail.Send(sender, m)
	}
	return nil

}
