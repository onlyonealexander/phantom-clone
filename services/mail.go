package services

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendQuestionEmail(question string) error {

	m := gomail.NewMessage()

	m.SetHeader(
		"From",
		os.Getenv("EMAIL_USER"),
	)

	m.SetHeader(
		"To",
		os.Getenv("ADMIN_EMAIL"),
	)

	m.SetHeader(
		"Subject",
		"New Question Received",
	)

	m.SetBody(
		"text/plain",
		question,
	)

	d := gomail.NewDialer(
    os.Getenv("SMTP_HOST"),
    587,
    os.Getenv("EMAIL_USER"),
    os.Getenv("EMAIL_PASS"),
)
	d.SSL = false
	d.TLSConfig = nil

	return d.DialAndSend(m)
}