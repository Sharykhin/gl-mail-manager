package main

import (
	"fmt"
	"log"

	"github.com/Sharykhin/gl-mail-manager/handler"
)

// {"action":"register", "payload":{"name":"serg", "to":"siarhei.sharykhin@itechart-group.com", "token":"12345"}}

func main() {
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	log.Fatal(handler.ListenAndServe())
}
