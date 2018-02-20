package gmail

import (
	"net/smtp"
	"os"
)

type SmtpGmailSender struct {
}

func (sm SmtpGmailSender) SendMail(from string, to string, subject string, body string) (string, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte("Subject:" + subject + "\n" + mime + body)
	from = os.Getenv("SMTP_GMAIL_USER")

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, os.Getenv("SMTP_GMAIL_PASS"), "smtp.gmail.com"),
		from, []string{to}, msg)

	return "gmailID", err
}
