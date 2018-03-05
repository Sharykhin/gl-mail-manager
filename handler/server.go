package handler

import (
	"fmt"
	"log"

	"github.com/Sharykhin/gl-mail-manager/queue/rabbitmq"
)

const (
	QUEUE_NAME = "mail"
)

// ListenAndServe listens income messages for a specific queue
func ListenAndServe() error {
	msgs, err := rabbitmq.Listen(QUEUE_NAME)
	if err != nil {
		return fmt.Errorf("could not listen %s queue: %v", QUEUE_NAME, err)
	}
	var done chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
			err := handle(d.Body)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	<-done
	return nil
}
