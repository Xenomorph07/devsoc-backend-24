package utils

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(to, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	dialer := gomail.NewDialer(
		os.Getenv("SMTP_SERVER"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_KEY"),
	)
	// dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} //need to be check later

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
