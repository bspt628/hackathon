package email

import (
	"fmt"
	"net/smtp"
)

type EmailSender struct {
	smtpHost string
	smtpPort string
	username string
	password string
}

func NewEmailSender(host, port, username, password string) *EmailSender {
	return &EmailSender{
		smtpHost: host,
		smtpPort: port,
		username: username,
		password: password,
	}
}

func (es *EmailSender) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", es.username, es.password, es.smtpHost)
	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))
	addr := fmt.Sprintf("%s:%s", es.smtpHost, es.smtpPort)
	return smtp.SendMail(addr, auth, es.username, []string{to}, msg)
}
