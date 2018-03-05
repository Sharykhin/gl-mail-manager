package mailgun

import (
	"os"

	"gopkg.in/mailgun/mailgun-go.v1"
)

// Provider is a reference to a private struct that implements all necessary methods
var Provider provider

type provider struct {
	mg mailgun.Mailgun
}

func (p provider) SendMail(from string, to string, subject string, body string) (string, error) {
	mgM := p.mg.NewMessage(
		from,
		subject,
		"",
		to,
	)
	mgM.SetHtml(body)
	_, id, err := p.mg.Send(mgM)
	return id, err
}

func init() {
	domain := os.Getenv("MG_DOMAIN")
	apiKey := os.Getenv("MG_API_KEY")
	publicKey := os.Getenv("MG_PUBLIC_API_KEY")
	mg := mailgun.NewMailgun(domain, apiKey, publicKey)
	Provider.mg = mg
}
