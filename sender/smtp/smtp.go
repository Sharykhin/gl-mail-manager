package gmail

import (
	"net/smtp"
	"os"
)

var (
	Provider provider
	login    string
	password string
)

type provider struct {
}

func (p provider) SendMail(from string, to string, subject string, body string) (string, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte("Subject:" + subject + "\n" + mime + body)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", login, password, "smtp.gmail.com"),
		from, []string{to}, msg)

	return "gmailID", err
}

func init() {
	login = os.Getenv("SMTP_GMAIL_USER")
	password = os.Getenv("SMTP_GMAIL_PASS")
}
