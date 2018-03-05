package handler

import (
	"fmt"

	"encoding/json"

	"github.com/Sharykhin/gl-mail-manager/entity"
	"github.com/Sharykhin/gl-mail-manager/grpc"
	"github.com/Sharykhin/gl-mail-manager/logger"
	"github.com/Sharykhin/gl-mail-manager/mail"
	"github.com/Sharykhin/gl-mail-manager/sender/mailgun"
)

const (
	TYPE_REGISTER = "register"
)

func handle(body []byte) error {
	mm := entity.MailMessage{}
	err := json.Unmarshal(body, &mm)
	if err != nil {
		return fmt.Errorf("could not parse income message: %v", err)
	}

	switch mm.Action {
	case TYPE_REGISTER:
		go sendRegisterMail(mm, body)
	default:
		fmt.Println("There is no action")
	}
	return nil
}

func sendRegisterMail(mm entity.MailMessage, body []byte) {
	fmt.Println("Send Register message")
	err := mail.SendRegisterMessage(mailgun.Provider, mm.Payload["to"].(string), mm.Payload)
	if err != nil {
		logger.Log.LogError("could not sent register email: " + string(body[:]))
		fmt.Printf("Could not sent register email: %s\n", err)
		_, err := grpc.CreateFailedMail(mm, err.Error())
		if err != nil {
			fmt.Printf("Could not create a new failed mail row on grpc server: %v", err)
		}
	}
}
