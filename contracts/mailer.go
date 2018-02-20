package contracts

type Mailer interface {
	SendMail(from string, to string, subject string, body string) (string, error)
}
