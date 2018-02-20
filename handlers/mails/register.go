package mails

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/Sharykhin/gl-mail-manager/contracts"
)

const MAIL_FROM = "Siarhei <siarhei.sharykhin@itechart-group.com>"

type registerPayload struct {
	Name  string
	Token string
}

// TODO: is it okay that we pass payload as map of interface values?
func SendRegisterMessage(m contracts.Mailer, to string, payload map[string]interface{}) error {
	// TODO: should we use register specific struct?
	p := registerPayload{
		Name:  payload["name"].(string),
		Token: payload["token"].(string),
	}

	dir, _ := os.Getwd()

	t, err := template.New("register.html").ParseFiles(dir + "/templates/register.html")
	if err != nil {
		return fmt.Errorf("could not parse template: %v", err)
	}

	var tpl bytes.Buffer
	t.Execute(&tpl, p)

	if os.Getenv("TEST_FAIL") == "OK" {
		log.Println("Test failed message")
		return err
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
