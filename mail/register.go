package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/Sharykhin/gl-mail-manager/contracts"
	"github.com/pkg/errors"
)

const MAIL_FROM = "Siarhei <siarhei.sharykhin@itechart-group.com>"

type registerPayload struct {
	Name  string
	Token string
}

var testFail string

// SendRegisterMessage send a register mail
func SendRegisterMessage(m contracts.Mailer, to string, payload map[string]interface{}) error {
	rp := registerPayload{
		Name:  payload["name"].(string),
		Token: payload["token"].(string),
	}

	t, err := template.New("register.html").ParseFiles("templates/register.html")
	if err != nil {
		return fmt.Errorf("could not parse template: %v", err)
	}

	var tpl bytes.Buffer
	t.Execute(&tpl, rp)

	if testFail == "OK" {
		fmt.Println("Test failed message")
		return errors.New("count not connect to mailgun")
	}

	id, err := m.SendMail(
		MAIL_FROM,
		to,
		"Welcome on board user",
		tpl.String(),
	)

	if err != nil {
		return fmt.Errorf("coulnd not sent mail: %v", err)
	}

	log.Println("ID: ", id)
	return nil
}

func init() {
	testFail = os.Getenv("TEST_FAIL")
}
