package messages

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Sharykhin/gl-mail-manager/contracts"
	"github.com/Sharykhin/gl-mail-manager/handlers/mails"
)

const TYPE_REGISTER = "register"

type MailMessage struct {
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}

func HandleMessage(body []byte, l contracts.Logger, mailer contracts.Mailer) error {
	mm := MailMessage{}
	err := json.Unmarshal(body, &mm)
	if err != nil {
		return fmt.Errorf("could not parse income message: %v", err)
	}

	switch mm.Action {
	case TYPE_REGISTER:
		log.Println("Send Register message")
		err := mails.SendRegisterMessage(mailer, mm.Payload["to"].(string), mm.Payload)
		if err != nil {
			l.LogError("could not sent register email: " + string(body[:]))
			log.Printf("Could not sent register email: %s\n", err)
		}
	default:
		log.Println("There is no action")
	}
	return nil
}
