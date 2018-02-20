package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

const ENV_ADDRESS = "AMPQ_ADDRESS"

func connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv(ENV_ADDRESS))
	if err != nil {
		return nil, fmt.Errorf("could not connect to rabbitmq: %v", err)
	}
	return conn, nil
}

func Listen(queue string) (<-chan amqp.Delivery, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("could not create a channel: %v", err)
	}

	notify := conn.NotifyClose(make(chan *amqp.Error))

	//TODO: Think whether this is the best place for listening broken connections?
	go listenClose(notify, ch, conn)

	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return nil, fmt.Errorf("could not decale queue %s: %v", queue, err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		true,   // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return nil, fmt.Errorf("could not create consume: %v", err)
	}

	return msgs, nil
}

func listenClose(notify chan *amqp.Error, ch *amqp.Channel, conn *amqp.Connection) {
	for {
		select {
		case err := <-notify:
			time.Sleep(time.Second * 5)
			ch.Close()
			conn.Close()
			log.Fatalf("rabbitmq connection was broken: %v", err)
		}
	}
}
