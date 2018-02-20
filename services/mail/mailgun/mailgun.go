package mailgun

import (
	"gopkg.in/mailgun/mailgun-go.v1"
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
	mg, err := mailgun.NewMailgunFromEnv()
	if err != nil {
		return nil, err
	}
	ms := &MailGunSender{instance: mg}
	return ms, nil
}
