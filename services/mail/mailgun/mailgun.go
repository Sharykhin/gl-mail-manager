package mailgun

import (
	"gopkg.in/mailgun/mailgun-go.v1"
	"os"
)

type MailGunSender struct {
	instance mailgun.Mailgun
}

func (ms MailGunSender) SendMail(from string, to string, subject string, body string) (string, error) {
	mgM := ms.instance.NewMessage(
		from,
		subject,
		"",
		to,
	)
	mgM.SetHtml(body)
	_, id, err := ms.instance.Send(mgM)
	return id, err
}

func New() (*MailGunSender, error) {
	//mg, err := mailgun.NewMailgunFromEnv()
	domain := os.Getenv("MG_DOMAIN")
	apiKey := os.Getenv("MG_API_KEY")
	publicKey := os.Getenv("MG_PUBLIC_API_KEY")
	mg := mailgun.NewMailgun(domain, apiKey, publicKey)
	ms := &MailGunSender{instance: mg}
	return ms, nil
}
