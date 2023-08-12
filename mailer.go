package emailer

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mailer struct {
	SmtpHost   string
	SmtpPort   int
	SenderName string
	Email      string
	Password   string
}

func New(senderName, email, password, smtpHost string, smtpPort int) *Mailer {
	return &Mailer{
		SmtpHost:   smtpHost,
		SmtpPort:   smtpPort,
		SenderName: senderName,
		Email:      email,
		Password:   password,
	}
}

func (m *Mailer) SendMailSync(to, cc []string, subject, message string) error {
	body := "MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"From: " + m.SenderName + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Cc: " + strings.Join(cc, ",") + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		message

	auth := smtp.PlainAuth("", m.Email, m.Password, m.SmtpHost)
	smtpAddr := fmt.Sprintf("%s:%d", m.SmtpHost, m.SmtpPort)

	err := smtp.SendMail(smtpAddr, auth, m.Email, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func (m *Mailer) SendMailAsync(to, cc []string, subject, message string) {
	go func() {
		body := "MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n" +
			"From: " + m.SenderName + "\r\n" +
			"To: " + strings.Join(to, ",") + "\r\n" +
			"Cc: " + strings.Join(cc, ",") + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			message

		auth := smtp.PlainAuth("", m.Email, m.Password, m.SmtpHost)
		smtpAddr := fmt.Sprintf("%s:%d", m.SmtpHost, m.SmtpPort)

		err := smtp.SendMail(smtpAddr, auth, m.Email, append(to, cc...), []byte(body))
		if err != nil {
			log.Println("Error sending email:", err)
		}
	}()
}
