package main

import (
	"log"

	"github.com/Sharykhin/gl-mail-manager/handlers/messages"
	"github.com/Sharykhin/gl-mail-manager/handlers/queue/rabbitmq"
	"github.com/Sharykhin/gl-mail-manager/providers"
	"github.com/Sharykhin/gl-mail-manager/services/mail/mailgun"
)

// {"action":"register", "payload":{"name":"serg", "to":"siarhei.sharykhin@itechart-group.com", "token":"12345"}}
func main() {

	msgs, err := rabbitmq.Listen("mail")
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	done := make(chan bool)
	go func() {
		l := providers.Logger()
		m, err := mailgun.New()
		if err != nil {
			log.Fatalf("counld not initalize mailgun: %v", err)
		}
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
			err := messages.HandleMessage(d.Body, l, m)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-done
}
